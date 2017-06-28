Vue.use(VueMaterial);
moment.locale('ja');

const app = new Vue({
  el: "#app",
  data: {
    tweets: [],
    users: [],
    tweet: {},
  },
  methods: {
    toggleUserList() {
      this.$refs.userList.toggle();
    },
    getAllTweets() {
      this.$http.get('/api/tweets').then(resp => {
        this.tweets = resp.body;
        this.$refs.userList.close();
      });
    },
    getUserTweets(userID) {
      this.$http.get(`/api/users/${userID}/tweets`).then(resp => {
        this.tweets = resp.body;
        this.$refs.userList.close();
      });
    },
    showTweetDetail(tweet) {
      if (tweet.text == null) {
        return;
      }
      this.tweet = tweet;
      this.$refs.tweetDetail.open();
    },
  },
  mounted() {
    this.getAllTweets();
    this.$http.get('/api/users').then(resp => this.users = resp.body);
  },
});
