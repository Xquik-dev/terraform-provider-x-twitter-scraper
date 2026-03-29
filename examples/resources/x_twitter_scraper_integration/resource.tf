resource "x-twitter-scraper_integration" "example_integration" {
  config = {
    chat_id = "chatId"
  }
  event_types = ["tweet.new"]
  name = "name"
  type = "telegram"
}
