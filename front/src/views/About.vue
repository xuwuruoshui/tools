<template>
    <div id="about">
     <div>pinia: {{ us.user.id }} --- {{ us.user.name }}</div>
      <button @click="change">修改</button>
      <button @click="change2">重置</button>
    </div>
</template>
<script setup lang='ts'>
import {userStore} from '@/stores'
const us = userStore()


// 监听state中的值变化
// us.$subscribe((args,state)=>{
//   console.log("=====>",args)
//   console.log("=====>",state)
// })

// 调用action时调用
us.$onAction((args) =>{
  args.after(()=>{
    console.log("执行完函数之后调用")
  })
  console.log(args)
})

let change = ()=>{
  us.setUser()
}
// 重置
let change2 = ()=>{
  us.$reset()
}
</script>

<style lang='scss' scoped>

</style>