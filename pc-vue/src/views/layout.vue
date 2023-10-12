<template>
  <div>
    <router-view></router-view>
    <Login />
  </div>
</template>

<script>
  import Login from '@/views/login/index'
  export default {
    name: 'layout',
    components: {Login}
  }
</script>
