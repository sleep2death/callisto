<template>
  <div class="home">
    <img
      alt="Vue logo"
      src="../assets/logo.png"
    >
    <HelloWorld :msg="this.$store.state.resp" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import axios from 'axios'

export default {
  name: 'home',
  components: {
    HelloWorld
  },
  mounted: function () {
    var code = this.$route.query.code
    console.log(code)
    axios.get(this.$store.state.AUTH_URL, { params: { code } })
      .then((resp) => {
        console.log(resp.data)
        this.$store.dispatch('response', resp.data['login'])
      })
  }
}
</script>
