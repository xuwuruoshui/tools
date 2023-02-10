<template>
  <div id="ws">
    <div id="content">
      <div id="record" v-for="(m,index) in respMsg" :key="index" :style="m.isOther?'color:red;':'color:blue;align-self:flex-end;'">
        {{m.data}}
      </div>
    </div>
    <input type="text" v-model="msg" id="msg-box">
    <input type="button" value="发送" @click="SendMessage">
  </div>

</template>

<script setup lang="ts">

import {reactive, ref} from "vue";

interface RespMsg{
  data: any
  isOther: boolean
}

let msg = ref<any>()
let respMsg = ref([])


let ws = new WebSocket("ws://192.168.2.39:3000/ws");
// 发送
let SendMessage = ()=>{
  let resp:RespMsg = {
    data: msg.value,
    isOther: false
  }
  ws.send(msg.value)
  respMsg.value?.push(resp)
  msg.value = ""
}

ws.onmessage =  function (e){
  let resp:RespMsg =  reactive({
    data: e.data,
    isOther: true
  })
  respMsg.value?.push(resp)
}

ws.onclose = function(e){
  //当客户端收到服务端发送的关闭连接请求时，触发onclose事件
  console.log("close");
}
ws.onerror = function(e){
  //如果出现连接、处理、接收、发送数据失败的时候触发onerror事件
  console.log(e);
}
</script>

<style lang="scss" scoped>
  #ws{
    width: 100%;
    height: 300px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  #content {
    width: 80%;
    height: 80%;
    border: 2px solid #000000;
  }
  #msg-box{
    width: 80%;
    height: 20%;
    margin: 5px;
  }
  #record{
    display: flex;
  }
</style>