resource "x-twitter-scraper_x_tweet" "example_x_tweet" {
  account = "@elonmusk"
  text = "Just launched our new feature!"
  attachment_url = "https://x.com/elonmusk/status/1234567890"
  community_id = "1500000000000000000"
  is_note_tweet = false
  media_ids = ["1234567890123456789"]
  reply_to_tweet_id = "1234567890"
}
