<template>
  <div class="app-container">
    <!--工具栏-->
    <div class="head-container">
      <div v-if="crud.props.searchToggle">
        <!-- 搜索 -->
        <el-select v-model="query.enabled" clearable size="small" placeholder="状态" class="filter-item" style="width: 90px" @change="crud.toQuery">
          <el-option v-for="item in enabledTypeOptions" :key="item.key" :label="item.display_name" :value="item.key" />
        </el-select>
        <rrOperation :crud="crud" />
      </div>
      <crudOperation :permission="permission" />
    </div>
    <!--表单组件-->
    <el-dialog append-to-body :close-on-click-modal="false" :before-close="crud.cancelCU" :visible.sync="crud.status.cu > 0" :title="crud.status.title" width="500px">
      <el-form ref="form" :model="form" :rules="rules" size="small" label-width="80px">
                    <el-form-item label="" prop="name">
                            <el-input v-model="form.name" placeholder=""/>
                    </el-form-item>
                    <el-form-item label="" prop="des">
                            <el-input v-model="form.des" placeholder=""/>
                    </el-form-item>
                    <el-form-item label="" prop="crateTime">
                            <el-date-picker
                                    v-model="form.crateTime"
                                    type="datetime"
                                    placeholder="选择日期">
                            </el-date-picker>
                    </el-form-item>
                    <el-form-item label="" prop="enable">
                            <el-input v-model="form.enable" placeholder=""/>
                    </el-form-item>
                    <el-form-item label="" prop="sort">
                            <el-input v-model="form.sort" placeholder=""/>
                    </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="text" @click="crud.cancelCU">取消</el-button>
        <el-button :loading="crud.cu === 2" type="primary" @click="crud.submitCU">确认</el-button>
      </div>
    </el-dialog>
    <!--表格渲染-->
    <el-table ref="table" v-loading="crud.loading" :tree-props="{children: 'children', hasChildren: 'hasChildren'}" default-expand-all :data="crud.data" row-key="id" @select="crud.selectChange" @select-all="crud.selectAllChange" @selection-change="crud.selectionChangeHandler">
      <el-table-column :selectable="checkboxT" type="selection" width="55" />
                <el-table-column label="" align="center" prop="id"
                                 :show-overflow-tooltip="true"/>
                <el-table-column label="" align="center" prop="name"
                                 :show-overflow-tooltip="true"/>
                <el-table-column label="" align="center" prop="des"
                                 :show-overflow-tooltip="true"/>
                <el-table-column label="" align="center" prop="crateTime"
                                 :show-overflow-tooltip="true">
                    <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.crateTime) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="" align="center" prop="updateTime"
                                 :show-overflow-tooltip="true">
                    <template slot-scope="scope">
                    <span>{{ parseTime(scope.row.updateTime) }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="" align="center" prop="enable"
                                 :show-overflow-tooltip="true"/>
                <el-table-column label="" align="center" prop="sort"
                                 :show-overflow-tooltip="true"/>
      <el-table-column v-permission="['admin','promptCategory:edit','promptCategory:del']" label="操作" width="130px" align="center" fixed="right">
        <template slot-scope="scope">
          <udOperation
            :data="scope.row"
            :permission="permission"
            :disabled-dle="scope.row.id === 1"
            msg="确定删除吗,如果存在下级节点则一并删除，此操作不能撤销！"
          />
        </template>
      </el-table-column>
    </el-table>
      <el-pagination
        :total="total"
        :current-page="page + 1"
        style="margin-top: 8px;"
        layout="total, prev, pager, next, sizes"
        @size-change="sizeChange"
        @current-change="pageChange"
      />
  </div>
</template>
<script>
import crudPromptCategory from '@/api/chatbot/PromptCategory'
import Treeselect from '@riophae/vue-treeselect'
import CRUD, { presenter, header, form, crud } from '@crud/crud'
import rrOperation from '@crud/RR.operation'
import crudOperation from '@crud/CRUD.operation'
import udOperation from '@crud/UD.operation'
// crud交由presenter持有
const defaultCrud = CRUD({ title: 'PromptCategory', url: 'admin/promptCategory', crudMethod: { ...crudPromptCategory }})
const defaultForm = {  id:"",  name:"",  des:"",  crateTime:"",  updateTime:"",  enable:"",  sort:"",    }
export default {
  name: 'PromptCategory',
  components: { Treeselect, crudOperation, rrOperation, udOperation },
  mixins: [presenter(defaultCrud), header(), form(defaultForm), crud()],
  // 设置数据字典
  dicts: ['dept_status'],
  data() {
    return {
      normalizer(node){
	        //去掉children=[]的children属性 respectively
	        if(node.children && !node.children.length){
	          delete node.children;
	        }
      },
      records: [],
      rules: {
        name: [
          { required: true, message: '请输入名称', trigger: 'blur' }
        ]
      },
      permission: {
        add: ['admin', 'promptCategory:add'],
        edit: ['admin', 'promptCategory:edit'],
        del: ['admin', 'promptCategory:del']
      },
      enabledTypeOptions: [
        { key: 1, display_name: '正常' },
        { key: 0, display_name: '禁用' }
      ]
    }
  },
  methods: {
    getPromptCategoryDatas(tree, treeNode, resolve) {
      const params = { pid: tree.id }
      setTimeout(() => {
        crudPromptCategory.getPromptCategory(params).then(res => {
          resolve(res.content)
        })
      }, 100)
    },
    // 新增与编辑前做的操作
    [CRUD.HOOK.afterToCU](crud, form) {
      if (form.pid !== null) {
        form.isTop = '0'
      } else if (form.id !== null) {
        form.isTop = '1'
      }
     // form.enabled = `${form.enabled}`
      if (form.id != null) {
        this.getSupPromptCategory(form.id)
      } else {
        this.getPromptCategory()
      }
    },
    getSupPromptCategory(id) {
      var data={"pid":id};
      crudPromptCategory.getPromptCategorySuperior(data).then(res => {
        const date = res.data.content
        this.buildPromptCategory(date)
        this.records = date
      })
    },
    buildPromptCategory(records) {
      (records || []).forEach(data => {
        if (data.children) {
          this.buildPromptCategory(data.children)
        }
        if (data.hasChildren && !data.children) {
          data.children = null
        }
      })
    },
    getPromptCategory() {
      crudPromptCategory.getPromptCategory({ enabled: 1 }).then(res => {
        this.records = res.data.content.map(function(obj) {
          if (obj.hasChildren) {
            obj.children = null
          }
          return obj
        })
      })
    },
    // 获取弹窗内部门数据
    loadPromptCategory({ action, parentNode, callback }) {
      if (action === LOAD_CHILDREN_OPTIONS) {
        crudPromptCategory.getPromptCategory({ enabled: true, pid: parentNode.id }).then(res => {
          parentNode.children = res.content.map(function(obj) {
            if (obj.hasChildren) {
              obj.children = null
            }
            return obj
          })
          setTimeout(() => {
            callback()
          }, 100)
        })
      }
    },
    // 提交前的验证
    [CRUD.HOOK.afterValidateCU]() {
      if (this.form.pid !== null && this.form.pid === this.form.id) {
        this.$message({
          message: '上级部门不能为空',
          type: 'warning'
        })
        return false
      }
      if (this.form.isTop === '1') {
        this.form.pid = null
      }
      return true
    },
    // 改变状态
    changeEnabled(data, val) {
      console.log("aa:"+JSON.stringify(this.dict.label.dept_status))
      console.log("aa:"+val)
      this.$confirm('此操作将 "' + this.dict.label.dept_status[val] + '" ' + data.name + '部门, 是否继续？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        crudPromptCategory.edit(data).then(res => {
          this.crud.notify(this.dict.label.dept_status[val] + '成功', CRUD.NOTIFICATION_TYPE.SUCCESS)
        }).catch(err => {
          data.enabled = !data.enabled
          console.log(err.response.data.message)
        })
      }).catch(() => {
        data.enabled = !data.enabled
      })
    },
    checkboxT(row, rowIndex) {
      return row.id !== 1
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>
::v-deep .vue-treeselect__control,::v-deep .vue-treeselect__placeholder,::v-deep .vue-treeselect__single-value {
  height: 30px;
  line-height: 30px;
}
</style>
<style rel="stylesheet/scss" lang="scss" scoped>
::v-deep .el-input-number .el-input__inner {
  text-align: left;
}
</style>
