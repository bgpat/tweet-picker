Vue.use(VueMaterial);

const app = new Vue({
  el: "#app",
  data: {
    tweets: [],
    users: [],
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
  },
  mounted() {
    this.getAllTweets();
    this.$http.get('/api/users').then(resp => this.users = resp.body);
  },
});
