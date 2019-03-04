<template>
  <div id="app">
    <transition>
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>

export default {
  name: 'app'
}
</script>

<style lang="scss">
html, body, #app {
  height: 100%;
}
</style>
