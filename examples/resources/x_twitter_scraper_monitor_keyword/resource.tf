resource "x-twitter-scraper_monitor_keyword" "example_monitor_keyword" {
  event_types = ["tweet.new"]
  query       = "xquik OR \"x api\""
}
