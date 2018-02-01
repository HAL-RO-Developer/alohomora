<template>
    <div>
        <h1>家のドアを開けてみよう笑笑</h1>
        <button class="button is-primary login-button" @click="onSubmit">開ける</button>
    </div>
</template>
<script>
    import api from "../../service/api"
    export default {
        name: "index",
        methods: {
            onSubmit: function (event) {
                var data = {
                    device_id:"apqxAep0dL6P",
                    port:[
                        {
                            "port_no": 1,
                            "function": 51,
                            "args":[1, 5000]
                        }
                    ]
                };
                api.func(data).then((res)=>{
                    console.log(res)
                })
            }
        }
    }
</script>
<style>
    .button{
        display: block;
        text-align: center;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 40%;
        height: 100px;
    }
    div{
        width: 100%;
        height: 100vh;
        background-color: aqua;
        position: relative;
    }
    h1{
        font-size: 1.4rem;
    }
</style>