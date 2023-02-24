import {defineStore} from 'pinia'
import {Stores} from './stores'

type User = {
    id: number
    name?: string
}

let result: User = {
    id: 1,
    name: "哈哈"
}

const Login = ():Promise<User> => {
    return new Promise(resolve => {
        setTimeout(() => {
            resolve({
                id: 2,
                name: "ffff"
            })
        }, 2000)
    })
}

export const userStore = defineStore(Stores.USER, {
    state: () => {
        return {
            user: <User>{id:0,name:"null"}
        }
    },
    // computed修饰一些值
    getters: {
        decroteUserName ():string{
            return "@@@"+this.user.name+"@@@"
        }
    },
    // methods 可以做同步 异步都可以做 提交state
    actions: {
        // 同步
        setUser() {
            this.user = result
        },
        // 异步
        async setUser2(){
            this.user = await Login()
            // 上一步执行完了才打印
            console.log("hha")
        }
    },
    persist: {
        // 设置存储哪里
        storage: localStorage
    }
})

