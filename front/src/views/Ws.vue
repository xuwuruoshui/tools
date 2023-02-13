<template>
  <div id="ws">
    <div id="content">
      <div id="record" v-for="(m,index) in respMsg" :key="index"
           :style="m.isOther?'color:red;':'color:blue;align-self:flex-end;'">
        {{ m.data }}
      </div>
    </div>
    <input type="text" v-model="msg" id="msg-box">
    <input type="button" value="发送" @click="SendMessage">
  </div>

</template>

<script setup lang="ts">

import {reactive, ref} from "vue";
import {ElMessage} from "element-plus";

interface RespMsg {
  data: string
  isOther: boolean
}

let msg = ref<string>()

let respMsg = ref<RespMsg[]>([])


let url: string = "ws://192.168.2.39:3000/ws"
let timeConnect = 0;

let ws:WebSocket
function webSocketInit(url:string) {
  ws = new WebSocket(url);
  ws.onopen = function () {
    console.log("已连接TCP服务器");
  };
  ws.onmessage = function (e) {
    let resp: RespMsg = reactive({
      data: e.data,
      isOther: true
    })
    respMsg.value?.push(resp)
  }
  ws.onclose = function (e) {
    //当客户端收到服务端发送的关闭连接请求时，触发onclose事件
    console.log("close");
    reconnect(url);
  }
  ws.onerror = function (e) {
    //如果出现连接、处理、接收、发送数据失败的时候触发onerror事件
    console.log(e);
  }
}
webSocketInit(url)
// 重连
function reconnect(url:string) {
  // lockReconnect加锁，防止onclose、onerror两次重连
  timeConnect++;
  console.log("第" + timeConnect + "次重连");
  // 进行重连
  setTimeout(function () {
    webSocketInit(url);
  }, 1000);

}



// 发送
let SendMessage = () => {
  if (msg.value == null) {
    ElMessage.error("内容不能为空!!!")
    return
  }
  let resp: RespMsg = {
    data: msg.value,
    isOther: false
  }

  ws.send(msg.value)
  respMsg.value?.push(resp)
  msg.value = ""
}

</script>

<style lang="scss" scoped>
#ws {
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
  display: flex;
  flex-direction: column;
}

#msg-box {
  width: 80%;
  height: 20%;
  margin: 5px;
}
</style>