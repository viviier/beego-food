<template>
  <div class="root">
    <h2>Beego-food</h2>
    <div class="signup">
      <el-input class="" v-model="name" placeholder="Name" clearable></el-input>
      <el-input class="" v-model="username" placeholder="Username" clearable></el-input>
      <el-input v-model="password" placeholder="Password" @enter="signup" type="password"></el-input>
      <router-link :to="{name: 'login'}">返回</router-link>
      <el-button type="primary" @click="signup">sign up</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'signup',
  data () {
    return {
      name: '',
      username: '',
      password: ''
    }
  },
  methods: {
    signup () {
      let _ = this
      axios({
        method: 'post',
        url: '/api/signup',
        data: {
          name: this.name,
          username: this.username,
          password: this.password
        }
      })
        .then(function (res) {
          console.log(res)
          _.$router.push({ name: 'foodlistall' })
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.root {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100%;

  h2 {
    margin-top: 0;
  }

  .signup {
    max-width: 360px;

    .el-input + .el-input {
      margin-top: 20px;
    }

    a {
      color: #03a9f4;
      display: flex;
      justify-content: flex-end;
      cursor: pointer;
      margin: 5px 0;
    }

    .el-button {
      width: 100%;
    }
  }
}
</style>
