<template>
  <div class="beego-header">
    <h2>Beego-food</h2>
    <el-menu :default-active="activeIndex" class="el-menu-demo" :router="true" mode="horizontal">
      <el-menu-item index="/foodlistall">首页</el-menu-item>
      <el-menu-item index="/shop">购物车</el-menu-item>
    </el-menu>
    <el-dropdown trigger="click" size="medium">
      <el-button type="primary" circle></el-button>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item>Login out</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
  </div>
</template>

<script>
export default {
  data () {
    return {
      activeIndex: ''
    }
  },
  created () {
    let pathname = document.location.pathname
    this.activeIndex = pathname
  }
}
</script>

<style lang="scss" scoped>
.beego-header {
  display: flex;

  h2 {
    font-size: 36px;
    margin-left: 40px;
    margin-right: 40px;
    cursor: default;
  }

  ul {
    flex: 1;
    align-self: center;
    border-bottom: none;

    li {
      margin-right: 20px;

      &:last-child {
        margin-right: 0;
      }
    }
  }

  .el-dropdown {
    height: 40px;
    align-self: center;
    margin-right: 40px;

    button {
      padding: 18px;
    }
  }
}
</style>
