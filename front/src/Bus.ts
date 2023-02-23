type BusClass = {
    emit:(name:string)=>void
    on:(name:string,callback:Function)=>void
}

type ParamsKey = string | number | symbol

type List = {
    // key为Params类型, value为函数数组
    [key:ParamsKey]:Array<Function>
}

class Bus implements BusClass{

    list:List

    constructor() {
        this.list = {}
    }

    emit(name:string,...args:Array<any>){
        let eventName:Array<Function>  = this.list[name]
        eventName.forEach(fn=>{
            // 调用函数
            fn.apply(this,args)
        })
    }

    on(name:string,callback:Function){
        let fn:Array<Function> = this.list[name] || []
        fn.push(callback)
        this.list[name] = fn
    }
}

export default new Bus()