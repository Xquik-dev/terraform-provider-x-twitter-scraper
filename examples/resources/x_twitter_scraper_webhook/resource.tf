resource "x-twitter-scraper_webhook" "example_webhook" {
  event_types = ["tweet.new", "follower.gained"]
  url = "https://example.com/webhook"
}
