<template>
  <div id="index">
    <!--#region 菜单路由 -->
    <el-row id="nav-menu" class="nav-menu" v-if="isShowSideBar">
      <el-menu :default-active="defaultActive" :default-openeds="['']"
               :collapse="isCollapse" :router="true"
               text-color="#78828a" active-text-color="#6777ef" background-color="#ffffff">
        <!--#region 折叠开关 -->
        <el-menu-item class="logo" index="">
          <el-row v-show="isCollapse">To</el-row>
          <template #title>
            <el-row>Tools</el-row>
          </template>
        </el-menu-item>
        <!--#endregion 折叠开关 -->
        <el-sub-menu index="">
          <template #title>
            <el-icon>
              <House/>
            </el-icon>
            <span>导航</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="/home" @click="collapseMenu2">
              <el-icon>
                <CoffeeCup/>
              </el-icon>
              Home
            </el-menu-item>
            <el-menu-item index="/about" @click="collapseMenu2">
              <el-icon>
                <More/>
              </el-icon>
              About
            </el-menu-item>
            <el-menu-item index="/ws" @click="collapseMenu2">
              <el-icon>
                <More/>
              </el-icon>
              Ws
            </el-menu-item>
            <el-menu-item index="/relation" @click="collapseMenu2">
              <el-icon>
                <More/>
              </el-icon>
              Relation
            </el-menu-item>
          </el-menu-item-group>
          <el-sub-menu index="1">
            <template #title>
              <el-icon>
                <location/>
              </el-icon>
              <span>用户管理</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/user">用户</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="/test">
          <el-icon>
            <document/>
          </el-icon>
          <template #title>Test</template>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon>
            <setting/>
          </el-icon>
          <template #title>Navigator Four</template>
        </el-menu-item>
      </el-menu>
    </el-row>
    <div id="nav-menu-cover" v-show="isShowCover" @click="collapseMenu3"></div>
    <!--#endregion 菜单路由 -->

    <!--#region 内容 -->
    <el-row id="content" justify="start">
      <el-col>
        <el-row id="header" justify="space-between">
          <el-icon :xs="8" :sm="4" :md="4" :lg="4" :xl="4" class="icon" @click="collapseMenu">
            <Expand v-show="!isCollapse"/>
            <Fold v-show="isCollapse"/>
          </el-icon>
          <el-col :xs="16" :sm="20" :md="20" :lg="20" :xl="20" class="userInfo">
            <el-row justify="end">
              <el-avatar class="avatar" :size="35"
                         src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"/>
              <el-dropdown class="userData">
            <span class="el-dropdown-link">
              你好，hahaha
              <el-icon class="el-icon--right">
                <arrow-down/>
              </el-icon>
            </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>Action 1</el-dropdown-item>
                    <el-dropdown-item>Action 2</el-dropdown-item>
                    <el-dropdown-item>Action 3</el-dropdown-item>
                    <el-dropdown-item disabled>Action 4</el-dropdown-item>
                    <el-dropdown-item divided>Action 5</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </el-row>
          </el-col>
        </el-row>
      </el-col>
      <el-col id="article">
        <section class="section">
          <div class="title">
            <h1>{{ title }} </h1>
          </div>
          <div class="content">
            <RouterView/>
          </div>
        </section>
      </el-col>
    </el-row>
    <!--#endregion 内容 -->

  </div>
</template>

<script setup lang="ts">
import 'element-plus/theme-chalk/display.css'
import {RouterView, useRouter} from 'vue-router'
import {computed, onMounted, Ref, ref, watch} from 'vue'
import {ArrowDown, CoffeeCup, Document, Expand, Fold, House, Location, More, Setting,} from '@element-plus/icons-vue'

// #region 折叠
const isCollapse = ref(true)
const sideFontSize = ref('1.1rem')
const collapseMenu = () => {
  isCollapse.value = !isCollapse.value

  if (isCollapse.value) {   // 折叠
    isShowSideBar.value = screenWidth.value >= 736;
    setTimeout(function () {
      sideFontSize.value = '2rem'
    }, 300)
  } else { // 打开
    isShowSideBar.value = true
    sideFontSize.value = '1.1rem'
  }
}


const collapseMenu2 = () => {
  if (screenWidth.value < 736) { // 小于736点击路由关闭侧边栏, 大于736不管
    isCollapse.value = !isCollapse.value
    if (isCollapse.value) {
      isShowSideBar.value = screenWidth.value >= 736;
      setTimeout(function () {
        sideFontSize.value = '2rem'
      }, 300)
    } else { // 打开
      isShowSideBar.value = true
      sideFontSize.value = '1.1rem'
    }
  }
}

const collapseMenu3 = () => {
  if (screenWidth.value < 736) { // 小于736点击路由关闭侧边栏, 大于736不管
    isCollapse.value = true
    isShowSideBar.value = false
    sideFontSize.value = '2rem'
  }
}
// #endregion

// #region 监听路由值的变化
let title = ref<string | symbol | null | undefined>("")
let defaultActive = ref<string>("/")
const router = useRouter()

watch(() =>
        router.currentRoute.value,
    (data) => {
      // 设置标题
      title.value = data.name
      // 设置导航值
      defaultActive.value = data.path

    }, {immediate: true, deep: true})
// #endregion

// #region 是否显示侧边栏
let isShowSideBar: Ref<Boolean> = ref<Boolean>(true)

// 监听屏幕宽度
window.onresize = () => {
  return (() => {
    screenWidth.value = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth
  })()
}
const screenWidth = ref(window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth)
watch(screenWidth, (width) => {
  if (width < 736) {
    isShowSideBar.value = false
  } else if (width > 736 && width <= 992) {
    isShowSideBar.value = true
    isCollapse.value = true
  }else if(width>992){
    isShowSideBar.value = true
    isCollapse.value = false
  }
})
// #endregion

// #region 遮罩层是否显示
let isShowCover = computed(() => {
  return screenWidth.value < 736 && isShowSideBar.value
})
// #endregion
onMounted(()=>{
  if (screenWidth.value < 736) {
    isShowSideBar.value = false
  } else if (screenWidth.value > 736 && screenWidth.value <= 992) {
    isShowSideBar.value = true
    isCollapse.value = true
  }else if(screenWidth.value>992){
    isShowSideBar.value = true
    isCollapse.value = false
  }
})
</script>

<style lang="scss" scoped>
#index {
  width: 100vw;
  height: 100vh;
  display: flex;
  position: relative;

  #nav-menu {
    height: 100%;


    .el-menu {
      border-right: 0;
      height: 100%;
    }

    // 侧边栏字体动态变化
    $sideFontSize: v-bind('sideFontSize');

    .el-sub-menu {
      font-size: $sideFontSize;

      .el-icon {
        font-size: $sideFontSize;
      }
    }

    .el-menu-item {
      font-size: $sideFontSize;

      .el-icon {
        font-size: $sideFontSize;
      }
    }

    .logo {
      color: #000000;
      justify-content: center;
      font-size: 1.1rem;
    }
  }

  @media screen and (max-width: 736px) {
    .nav-menu {
      position: absolute;
      z-index: 2;
    }
  }

  #nav-menu-cover {
    position: absolute;
    z-index: 1;
    width: 100%;
    height: 100%;
    background: #000000;
    opacity: 0.5;
  }

  #content {
    flex-grow: 1;
    height: 100%;
    background: #f4f6f9;
    align-content: flex-start;


    #header {
      background: #6777ef;
      cursor: pointer;
      font-size: 1.5rem;
      align-items: center;

      .icon {
        color: #f2f2f2;
        margin: 20px;
      }

      .userInfo {
        margin: 20px;
        font-size: 1rem;

        .avatar {
        }

        .userData {
          color: #fefeff;
          margin: 10px;
        }
      }
    }

    #article {
      width: 100%;
      display: flex;
      flex-direction: column;
      align-items: center;
      color: #34395e;


      .section {
        width: 95%;


        .title {
          background-color: #ffffff;
          padding: 10px;
          border-radius: 5px;
          //position: relative;
          //left: 0px;
          //top: -20px;
          box-shadow: rgb(0 0 0) 0px 2px 8px -7px;

          h1 {
            padding: 20px;
          }
        }

        .content {
          background-color: #ffffff;
          padding: 30px;
          position: relative;
          left: 0;
          top: -20px;
          margin-top: 10px;
          border-radius: 5px;
          box-shadow: rgb(0 0 0) -1px 1px 8px -7px;
        }
      }


    }

  }
}
</style>