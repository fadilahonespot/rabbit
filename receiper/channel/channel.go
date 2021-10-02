package channel

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Channel1(c *amqp.Delivery) (err error) {
	if c.Body == nil {
		fmt.Println("Error")
		return fmt.Errorf("value is nil")
	}
	fmt.Println("THIS CHANNEL 1: ", string(c.Body))
	return nil
}

func Channel2(c *amqp.Delivery) (err error) {
	if c.Body == nil {
		fmt.Println("Error")
		return fmt.Errorf("value is nil")
	}
	fmt.Println("THIS CHANNEL 2: ", string(c.Body))
	return nil
}

func Channel3(c *amqp.Delivery) (err error) {
	if c.Body == nil {
		fmt.Println("Error")
		return fmt.Errorf("value is nil")
	}
	fmt.Println("THIS CHANNEL 3: ", string(c.Body))
	return nil
}

func Channel4(c *amqp.Delivery) (err error) {
	if c.Body == nil {
		fmt.Println("Error")
		return fmt.Errorf("value is nil")
	}
	fmt.Println("THIS CHANNEL 4: ", string(c.Body))
	return nil
}

func Channel5(c *amqp.Delivery) (err error) {
	if c.Body == nil {
		fmt.Println("Error")
		return fmt.Errorf("value is nil")
	}
	fmt.Println("THIS CHANNEL 5: ", string(c.Body))
	return nil
}