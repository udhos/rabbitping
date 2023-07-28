package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func pinger(app *application) {
	const me = "pinger"

	var countErrors int

	for {
		begin := time.Now()
		_, errDial := amqpDial(app.conf.amqpURL, app.conf.timeout)
		elap := time.Since(begin)

		var outcome string
		if errDial != nil {
			countErrors++
			outcome = "error"
			log.Printf("%s: elapsed=%v outcome=%s errors=%d/%d error:%v",
				me, elap, outcome, countErrors, app.conf.failureThreshold, errDial)
		} else {
			countErrors = 0
			outcome = "success"
			log.Printf("%s: elapsed=%v outcome=%s errors=%d/%d",
				me, elap, outcome, countErrors, app.conf.failureThreshold)
		}
		app.met.recordLatency(outcome, elap)

		if countErrors >= app.conf.failureThreshold {
			countErrors = 0
			if app.conf.restartDeploy != "" {
				restartDeploy(app.conf.restartNamespace, app.conf.restartDeploy)
			}
		}

		log.Printf("%s: sleeping for %v", me, app.conf.interval)
		time.Sleep(app.conf.interval)
	}
}

func restartDeploy(namespace, deployment string) {
	const me = "action"

	log.Printf("%s: failure threshold violated, restart namespace=%s deploy=%s",
		me, namespace, deployment)

	config, errConfig := rest.InClusterConfig()
	if errConfig != nil {
		log.Printf("%s: running OUT-OF-CLUSTER: %v", me, errConfig)
		return
	}

	clientset, errClientset := kubernetes.NewForConfig(config)
	if errClientset != nil {
		log.Printf("%s: kube clientset error: %v", me, errClientset)
		return
	}

	ctx := context.TODO()

	deploy, errGetDeploy := clientset.AppsV1().Deployments(namespace).Get(ctx, deployment, metav1.GetOptions{})
	if errGetDeploy != nil {
		log.Printf("%s: kube get deploy error: %v", me, errGetDeploy)
		return
	}

	if deploy.Spec.Template.ObjectMeta.Annotations == nil {
		deploy.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}

	deploy.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	_, errUpdate := clientset.AppsV1().Deployments(namespace).Update(ctx, deploy, metav1.UpdateOptions{})
	if errUpdate != nil {
		log.Printf("%s: kube update error: %v", me, errUpdate)
		return
	}

	log.Printf("%s: failure threshold violated, restart namespace=%s deploy=%s: done",
		me, namespace, deployment)
}

// amqp.Dial but with timeout
func amqpDial(amqpURL string, timeout time.Duration) (*amqp.Connection, error) {
	//conn, err := amqp.Dial(amqpURL)
	config := amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
		Dial:      amqp.DefaultDial(timeout),
	}
	conn, err := amqp.DialConfig(amqpURL, config)
	return conn, err
}
