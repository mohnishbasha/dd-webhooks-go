package main

import "fmt"

var currentId int

var webhooks Webhooks

// Give us some seed data
func init() {
	RepoCreateWebhook(Webhook{Name: "Write presentation"})
	RepoCreateWebhook(Webhook{Name: "Host meetup"})
}

func RepoFindWebhook(id int) Webhook {
	for _, t := range webhooks {
		if t.Id == id {
			return t
		}
	}
	// return empty Webhook if not found
	return Webhook{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateWebhook(t Webhook) Webhook {
	currentId += 1
	t.Id = currentId
	webhooks = append(webhooks, t)
	return t
}

func RepoDestroyWebhook(id int) error {
	for i, t := range webhooks {
		if t.Id == id {
			webhooks = append(webhooks[:i], webhooks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Webhook with id of %d to delete", id)
}
