resource "x-twitter-scraper_x_tweet" "example_x_tweet" {
  account = "@elonmusk"
  community_id = "1500000000000000000"
  is_note_tweet = false
  media = ["https://example.com/video.mp4"]
  reply_to_tweet_id = "1234567890"
  text              = "Just launched our new feature!"
}
