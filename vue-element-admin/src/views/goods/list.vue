<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="转链" width="140">
        <template slot-scope="scope">
          <router-link :to="'/goods/genwxlink/'+scope.row.id">
            <el-button type="primary" size="small">
              微信
            </el-button>
          </router-link>
          <router-link :to="'/goods/genqqlink/'+scope.row.id">
            <el-button type="primary" size="small">
              QQ
            </el-button>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="图片" width="110">
        <template slot-scope="{row}">
          <a target="_blank" :href=row.itemLink >
            <img :src="row.mainPic" width="80" height="80">
          </a>
        </template>   
      </el-table-column>
      <el-table-column align="center" label="标题" width="300">
        <template slot-scope="scope">
          <span>{{ scope.row.dtitle }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="淘宝商品ID" width="115">
        <template slot-scope="scope">
            <span>{{ scope.row.goodsId }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="大淘客ID" width="90">
        <template slot-scope="scope">
          <span>{{ scope.row.dataokeId }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="原价" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.originalPrice }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="折扣价" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.actualPrice }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="佣金" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.actualPrice * scope.row.commissionRate / 100 }} ({{scope.row.commissionRate}}%)</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="优惠券" width="100">
        <template slot-scope="{row}">
          <a target="_blank" :href=row.couponLink >
            满{{ row.couponConditions }}减{{ row.couponPrice }}
          </a>
        </template>  
      </el-table-column>
      <el-table-column align="center" label="淘宝cid" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.tbcid }}</span>
        </template>
      </el-table-column>

    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList } from '@/api/goods'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import splitPane from 'vue-splitpane'

export default {
  name: 'GoodsList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.list = response.data.items
        this.total = response.data.total
        this.listLoading = false
      })
    }
  }
}

</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
