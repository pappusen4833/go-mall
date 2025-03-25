/**
    腾讯地图组件

 */
<template>
    <div class="tencentmap">
        <div id="container"></div>
        <ul class="addressBtn">
            <li class="btnContainer" v-for="(item, index) in storeList" :key="index" @click="changeAddress(storeList[index]),cssActive(index)" :class="{active_Container : index === storeIndex.index}">
                <div class="title">{{ item.name }}</div>
                <div class="address">
                    <el-image class="mar-right-5" src="../../../static/images/zb.png" ></el-image>{{ item.address + item.detailedAddress }}
                </div>
                <div class="workTime">
                    <span class="time"><el-image class="mar-right-5" src="../../../static/images/time.png" ></el-image>{{ item.dayTime }}</span>
                <!-- <span class="distance">距离{{item.distance}}m</span> -->
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    props: {
        storeList: Array,
        storeIndex: Object
    },
    mounted () {
        console.log('挂载')
        const oScript = document.createElement('script')
        oScript.type = 'text/javascript'
        oScript.src = 'https://map.qq.com/api/gljs?v=1.exp&key=OGABZ-Y5OCF-5UWJ5-N7DHH-VFIG7-DHFEB'
        document.body.appendChild(oScript)
        if (oScript.readyState) {
            oScript.onreadyStateChange = () => {
                if (oScript.readyState === 'complete' || oScript.readyState === 'loaded') {
                     this.initMap(this.storeList[this.storeIndex.index])
                }
            }
        } else {
            oScript.onload = () => {
                this.initMap(this.storeList[this.storeIndex.index])
            }
        }
    },
    data () {
        return {
            isInit: false,
            map: null,
            markerLayer: null
        }
    },
    methods: {
        initMap (x) {
            /* eslint-disable */
            var center = new TMap.LatLng(x.latitude, x.longitude)
            this.map = new TMap.Map(document.getElementById('container'), {
                center: center,
                zoom: 15
            })
            console.log('初始化')
            var geometries = this.storeList.map((item, index) => {
                return {
                    'id': `${index}`,
                    'styleId': 'myStyle',
                    'position': new TMap.LatLng(item.latitude, item.longitude)
                }
            })
            this.markerLayer = new TMap.MultiMarker({
                map: this.map,
                styles: {
                    'myStyle': new TMap.MarkerStyle({
                        'width': 25,
                        'height': 35,
                        'src': '../../static/images/dwicon.png',
                        'anchor': { x: 16, y: 32 }
                    })
                },
                geometries: geometries
            })
            this.isInit = true
        },
        changeAddress (val) {
            if (!this.initMap) {
                console.log('地图未初始化完毕')
                return
            }
            this.map.setCenter(new TMap.LatLng(val.latitude, val.longitude))
        },
        cssActive (index) {
            this.storeIndex.index = index
        },
        geocoder () {
            /**
             * url：https://apis.map.qq.com/ws/geocoder/v1
             * methods: get
             * params: {
             *  "address": "深圳市南山区海德二道2748号茂业百货7楼",
             *  "key": "OGABZ-Y5OCF-5UWJ5-N7DHH-VFIG7-DHFEB"
             * }
             * 响应示例：
             {
                "status": 0, //0正常，其他皆为异常，具体查看message
                "message": "query ok",
                "result": {
                    "title": "茂业百货(南山店)",
                    "location": {
                        "lng": 113.930367,
                        "lat": 22.519361
                    },
                    "ad_info": { //行政区划信息
                        "adcode": "440305" //行政区划代码
                    },
                    "address_components": {
                        "province": "广东省",
                        "city": "深圳市",
                        "district": "南山区",
                        "street": "海德二道",
                        "street_number": "2748"
                    },
                    "reliability": 7, //可信度参考：值范围 1 <低可信> - 10 <高可信>
                    "level": 11 //解析精度级别，分为11个级别，一般>=9即可采用（定位到点，精度较高）
                }
            }
            */
        }
    }
}
</script>

<style lang="scss">
.tencentmap{
    .addressBtn{
        .btnContainer{
            width: 670px;
            height: 135px;
            margin: 20px auto;
            padding: 20px;
            cursor: pointer;
            border: 1px solid #CCCCCC;
            .title{
                font-size: 24px;
                color: #333333;
                font-weight: 400;
            }
            .address{
                display: flex;
                align-items: center;
                margin: 15px 0;
                img{
                    width: 15px;
                    height: 15px;
                }
            }
            .workTime{
                display: flex;
                justify-content: space-between;
                .time{
                    display: flex;
                    align-items: center;
                    img{
                        width: 15px;
                        height: 15px;
                    }
                }
            }
        }
        .active_Container{
            background-color: rgba($color: #F26E47, $alpha: .1);
            border: 1px solid #F26E47;
        }
    }
}
</style>
