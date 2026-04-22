resource "x-twitter-scraper_webhook" "example_webhook" {
  event_types = ["tweet.new", "tweet.reply"]
  url = "https://example.com/webhook"
}
