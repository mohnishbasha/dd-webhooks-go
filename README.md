# dd-webhooks-go
Simple WebApplication in GO - POST Handler

REST URL
http://52.38.25.212:8080/webhooks

POST Data

{
"eventid":$ID,
"orgid":$ORG_ID,
"title":"$EVENT_TITLE", 
"eventtype":"$EVENT_TYPE",
"eventbody":"$EVENT_MSG",
"name":"$ORG_NAME"
}

