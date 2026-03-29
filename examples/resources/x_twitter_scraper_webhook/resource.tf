resource "x-twitter-scraper_webhook" "example_webhook" {
  event_types = ["tweet.new"]
  url = "https://example.com"
}
