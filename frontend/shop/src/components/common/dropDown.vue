<template>
    <div class="dropDown">
        <div class="dropDown_select"
        :style="{width:`${width}px`,height:`${height}px`}">
            <div class="dropDown_title">
                {{ this.selectTitle }}<i class="el-icon-arrow-down el-icon--right"></i>
            </div>
            <div class="dropDown_content">
                <ul>
                    <li
                    class="dropItem"
                    :style="{width:`${width}px`,height:`${height}px`}"
                    v-for="(item,index) in arr"
                    :key="index"
                    @click="handleChange(item)">{{ item.label }}</li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        width: {
            typeof: Number,
            default: 100
        },
        height: Number,
        arr: {
            typeof: Array,
            default: () => {
                return []
            }
        },
        title: {
            typeof: String,
            default: '默认'
        }
    },
    data () {
        return {
            selectTitle: ''
        }
    },
    created () {
        this.selectTitle = this.title
    },
    methods: {
        handleChange (item) {
            this.selectTitle = item.label
            this.$emit('selectVal', item)
        }
    }
}
</script>

<style lang="scss">
.dropDown{
    .dropDown_select{
        width: 100px;
        height: 50px;
        overflow: hidden;
        margin: auto;
        .dropDown_title{
            width: 100%;
            text-align: center;
            border-radius: 3px;
            cursor: pointer;
            i{
                float: right;
                line-height: 50px;
                margin-right: 5px;
            }
        }
        .dropDown_content{
            position: relative;
            z-index: 1;
            background-color: white;
        }
        ul{
            // border: 1px solid #000;
            border-radius: 3px;
            box-shadow: 0px 0px 3px rgba(144, 144, 153, .6);
        }
        .dropItem{
            width: 100px;
            height: 40px;
            line-height: 30px;
            padding: 5px 10px;
            text-align: center;
            cursor: pointer;
            // color: #409EFF;
            // border: 1px solid #409EFF;
            border-radius: 5px;
            &:hover{
                // color: #FFF;
                background-color: rgba(144, 144, 153, .3);
            }
        }
        &:hover{
            // overflow: inherit;//从其父元素继承overflow值,默认值visible
            overflow: visible;
            box-shadow: none;
        }
        &:hover .dropDown_content{
            animation: opa 1s;
            position: absolute;
        }
    }
}
@keyframes opa {
    from{
        opacity: 0;
    }
    to{
        opacity: 1;
    }
}
</style>
