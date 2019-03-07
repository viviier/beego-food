<template>
  <div class="root">
    <h2>Beego-food</h2>
    <div class="login">
      <el-input class="" v-model="username" placeholder="Username" clearable></el-input>
      <el-input v-model="password" placeholder="Password" @keyup.enter.native="login" type="password"></el-input>
      <router-link :to="{ name: 'signup' }">create new account</router-link>
      <el-button type="primary" @click="login">login</el-button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'login',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    login () {
      let _ = this
      axios({
        method: 'post',
        url: '/api/login',
        data: {
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

  .login {
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
