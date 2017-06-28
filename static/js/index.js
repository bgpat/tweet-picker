Vue.use(VueMaterial);
moment.locale('ja');

const app = new Vue({
  el: "#app",
  data: {
    tweets: [],
    users: [],
    tweet: {},
  },
  computed: {
    parsedTweet() {
      if (this.tweet == null || this.tweet.entities == null) {
        return;
      }
      let entities = [];
      for (let key in this.tweet.entities) {
        if (!Array.isArray(this.tweet.entities[key])) {
          continue;
        }
        this.tweet.entities[key].forEach(entity => {
          entity.key = key;
          entities.push(entity);
        });
      }
      entities = entities.sort((a, b) => a.indices[0] - b.indices[0]);
      let tweet = [this.tweet.text];
      let i = 0;
      for (let e of entities) {
        let begin = e.indices[0] - i;
        let end = e.indices[1] - i;
        let last = tweet.pop();
        tweet.push(last.slice(0, begin), e, last.slice(end));
        i += end;
      }
      return tweet;
    },
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
      console.log(JSON.parse(JSON.stringify(tweet)));
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
