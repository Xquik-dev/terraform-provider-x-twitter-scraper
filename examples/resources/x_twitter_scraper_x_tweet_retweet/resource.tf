resource "x-twitter-scraper_x_tweet_retweet" "example_x_tweet_retweet" {
  account         = "@example"
  idempotency_key = "terraform-retweet-v1"
  target_id       = "1234567890"
}
