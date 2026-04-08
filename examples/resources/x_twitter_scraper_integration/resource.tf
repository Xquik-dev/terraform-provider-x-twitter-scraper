resource "x-twitter-scraper_integration" "example_integration" {
  config = {
    chat_id = "-1001234567890"
  }
  event_types = ["tweet.new", "follower.gained"]
  name = "My Telegram Bot"
  type = "telegram"
}
