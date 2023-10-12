<template>
  <div class="selectBox">
    <div class="noTypeList" v-if="this.data.length <= 0">
      暂无数据~
    </div>
    <ul class="typeList" v-if="this.data.length > 0">
      <li class="liItem"
        v-for="item in this.data" :key="item.value"
        :class="{liActive : typeSelectItem === item.value}"
        @click="getFatherValue(item)">
        {{ item.label }}
        <i class="el-icon-arrow-right" />
      </li>
    </ul>
    <ul class="typeItemContainer" v-if="typeSon.length > 0">
      <li class="sonItem"
        v-for="item in this.typeSon"
        :key="item.id"
        @click="selectTypeItem(item)"
        :title="item.label">
        <el-image class="labelIcon" :src="item.pic" ></el-image>
        <span>{{ item.label}}</span>
      </li>
    </ul>
    <div class="typeItemContainer NoChildren" v-if="typeSon.length <= 0">
      <NotFoundInfo />
    </div>
  </div>
</template>

<script>
import NotFoundInfo from '@/components/NotFound/NotFoundInfo.vue'
export default {
  components: {
    NotFoundInfo
  },
  props: {
    data: {
      typeof: Array,
      default: []
    },
    typeSelect: {
      typeof: String,
      default: ''
    }
  },
  data () {
    return {
      typeSelectItem: this.typeSelect, // 父级id
      typeSon: []
    }
  },
  watch: {
    typeSelect: {
      handler (nVal, oVal) {
        this.typeSelectItem = nVal
        this.data.forEach(item => {
          if (item.value === nVal) {
            this.typeSon = item.children
          }
        })
      }
    }
  },
  methods: {
    getFatherValue (item) {
      console.log(item)
      this.typeSelectItem = item.value
      this.typeSon = item.children
    },
    selectTypeItem (item) {
      console.log(item)
      this.$router.push({
        path: '/indexpage/productlist',
        query: {
          type1: this.typeSelectItem,
          type2: item.value || ''
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.selectBox{
  &:hover .typeItemContainer{
    display: inline-block;
    animation: changeWidth .5s;
  }
  .typeList{
    width: 250px;
    background-color: rgba($color: #000000, $alpha: .3);
    .liItem{
      height: 56px;
      line-height: 56px;
      padding: 0 30px;
      color: #FFFFFF;
      font-size: 16px;
      font-family: PingFang SC;
      font-weight: bold;
      cursor: pointer;
      .el-icon-arrow-right{
        float: right;
        line-height: 56px;
      }
      &:hover{
        background-color: #000000;
        display: flex;
        flex-wrap: wrap;
      }
    }
    .liActive{
      background-color: #000000;
    }
  }
  .typeItemContainer{
    display: none;
    width: 718px;
    height: 100%;
    padding: 12px 8px;
    background-color: #FFFFFF;
    .sonItem{
      display: inline-flex;
      width: 20%;
      margin: 0 2%;
      height: 70px;
      align-items: center;
      cursor: pointer;
      border-radius: 5px;
      &:hover {
        background-color: #F1F1F1;
      }
      .labelIcon{
        width: 56px;
        height: 56px;
        margin: 14px;
      }
      span{
        width: 100%;
        font-size: 16px;
        font-family: PingFang SC;
        font-weight: bold;
        color: #333333;
        overflow: hidden;
        text-overflow: ellipsis;
      }
    }
  }
  .noTypeList{
    width: 250px;
    background-color: rgba($color: #000000, $alpha: .3);
    display: flex;
    justify-content: center;
    align-items: center;
    color: #FFF;
    font-weight: bold;
  }
}
@keyframes changeWidth {
  from{
      opacity: 0;
  }
  to {
      opacity: 1;
  }
}
</style>
