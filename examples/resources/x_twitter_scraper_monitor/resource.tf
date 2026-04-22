resource "x-twitter-scraper_monitor" "example_monitor" {
  event_types = ["tweet.new", "tweet.reply"]
  username = "elonmusk"
}
