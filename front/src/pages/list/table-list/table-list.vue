<template>
  <div>
    <Button type="primary" icon="md-add" @click="handleOpenCreate">New</Button>
    <div class="ivu-inline-block ivu-fr">
      <Dropdown @on-click="handleChangeTableSize" trigger="click">
        <Tooltip class="ivu-ml" content="" placement="top">
          <i-link>
            <Icon type="md-list"/>
          </i-link>
        </Tooltip>
        <DropdownMenu slot="list">
          <DropdownItem name="default" :selected="tableSize === 'default'">Default</DropdownItem>
          <DropdownItem name="large" :selected="tableSize === 'large'">Large</DropdownItem>
          <DropdownItem name="small" :selected="tableSize === 'small'">Small</DropdownItem>
        </DropdownMenu>
      </Dropdown>
      <Tooltip class="ivu-ml" :content="tableFullscreen ? 'Quit Fullscreen' : 'Fullscreen'" placement="top">
        <i-link @click.native="handleFullscreen">
          <Icon :custom="tableFullscreen ? 'i-icon i-icon-exit-full-screen' : 'i-icon i-icon-full-screen'"/>
        </i-link>
      </Tooltip>
      <Tooltip class="ivu-ml" content="Refresh" placement="top">
        <i-link @click.native="handleRefresh">
          <Icon custom="i-icon i-icon-refresh"/>
        </i-link>
      </Tooltip>
      <Dropdown trigger="click">
        <Tooltip class="ivu-ml" content="Column Setting" placement="top">
          <i-link>
            <Icon type="md-options"/>
          </i-link>
        </Tooltip>
        <DropdownMenu slot="list">
          <div class="ivu-p-8">
            <Row>
              <Col span="12">Columns</Col>
              <Col span="12" class="ivu-text-right">
                <i-link link-color @click.native="handleResetColumn">Refresh</i-link>
              </Col>
            </Row>
          </div>
          <Divider size="small" class="ivu-mt-8 ivu-mb-8"/>
          <li class="ivu-dropdown-item" v-for="item in columns" :key="item.title" v-if="item.title"
              @click="item.show = !item.show">
            <Checkbox v-model="item.show"></Checkbox>
            <span>{{ item.title }}</span>
          </li>
        </DropdownMenu>
      </Dropdown>
    </div>
    <Alert show-icon class="ivu-mt" v-if="untouched !== 0">
      <div v-font="14">
        Have <strong v-color="'#2d8cf0'">{{ untouched }}</strong> Patients Not Contact yet
      </div>
    </Alert>
    <Alert type="success" show-icon class="ivu-mt" v-if="untouched === 0 && total !== 0">
      <div v-font="14">
        All Patients Have Been Contacted, Good Job!
      </div>
    </Alert>
    <Table
        ref="table"
        :columns="tableColumns"
        :data="list"
        :loading="loading"
        :size="tableSize"
        class="ivu-mt"
        @on-sort-change="handleSortChange"
        @on-filter-change="handleFilterChange"
        @on-select="handleSelect"
        @on-select-cancel="handleSelectCancel"
        @on-select-all="handleSelectAll"
        @on-select-all-cancel="handleSelectAllCancel"
    >
      <template slot-scope="{ row }" slot="Status">
        <Badge v-if="row.Status === 0" status="processing" text="Waiting"/>
        <Badge v-if="row.Status === 1" status="success" text="Contacted"/>
      </template>
      <template slot-scope="{ row }" slot="Appointment">
        {{ row.Appointment | dealWithTz(1) }}
      </template>
      <template slot-scope="{ row }" slot="Birthday">
        {{ row.Birthday | dealWithTz(0) }}
      </template>
      <template slot-scope="{ row }" slot="CreatedAt">
        {{ row.CreatedAt | dealWithTz(1) }}
      </template>
      <template slot-scope="{ row }" slot="UpdatedAt">
        {{ row.UpdatedAt | dealWithTz(1) }}
      </template>
      <template slot-scope="{ row }" slot="Photo">
        <a @click="handleClickView(row.ID)">Click To View</a>
      </template>
      <template slot-scope="{ row, index }" slot="action">
        <a @click="handleUpdate(row.ID)" v-if="row.Status === 0">Contacted</a>
      </template>
    </Table>
    <div class="ivu-mt ivu-text-center">
      <Page
          :total="total"
          :current.sync="current"
          show-total
          show-sizer
          :page-size="size"
          @on-page-size-change="handleChangePageSize"
          @on-change="handleChangePage"
      />
    </div>

    <Modal v-model="showCreate" title="Edit" :loading="creating" @on-ok="handleCreate" :transfer="false">
      <Form ref="create" :model="createData" :rules="createRules" :label-width="80">
        <FormItem label="Desc：" prop="desc">
          <Input v-model="createData.desc" placeholder=""/>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>
<script>
    import { GetUserLists, UpdateUserStatus } from '@api/users';

    export default {
        data () {
            return {
                columns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center',
                        show: true
                    },
                    {
                        title: 'Name',
                        key: 'Name',
                        minWidth: 120,
                        show: true
                    },
                    {
                        title: 'Birthday',
                        slot: 'Birthday',
                        minWidth: 120,
                        show: true
                    },
                    {
                        title: 'Phone',
                        key: 'Phone',
                        minWidth: 120,
                        show: true
                    },
                    {
                        title: 'Email',
                        key: 'Email',
                        minWidth: 160,
                        show: true
                    },
                    {
                        title: 'Appointment',
                        slot: 'Appointment',
                        minWidth: 140,
                        show: true
                    },
                    {
                        title: 'Status',
                        slot: 'Status',
                        minWidth: 100,
                        show: true
                    },
                    {
                        title: 'Photo',
                        slot: 'Photo',
                        minWidth: 100,
                        show: true
                    },
                    {
                        title: 'Address',
                        key: 'Address',
                        minWidth: 180,
                        show: true
                    },
                    {
                        title: 'CreatedAt',
                        slot: 'CreatedAt',
                        minWidth: 140,
                        sortable: 'custom',
                        show: true
                    },
                    {
                        title: 'UpdatedAt',
                        slot: 'UpdatedAt',
                        minWidth: 140,
                        sortable: 'custom',
                        show: true
                    },
                    {
                        title: 'Operate',
                        slot: 'action',
                        align: 'center',
                        minWidth: 140,
                        show: true
                    }
                ],
                loading: false,
                list: [],
                selectedData: [],
                current: 1,
                size: 20,
                sortType: 'desc', //  asc/desc
                sortColumns: 'id',
                filterType: undefined,
                showCreate: false,
                createData: {
                    desc: ''
                },
                createRules: {
                    desc: [
                        { required: true, message: 'DESC', trigger: 'blur' }
                    ]
                },
                creating: true,
                updateIndex: -1,
                tableSize: 'default',
                tableFullscreen: false,
                total: 0,
                untouched: 0
            }
        },
        filters: {
            dealWithTz: function (value, type) {
                if (!value) return ''
                let val = new Date(value)
                let options
                if (type === 1) {
                    options = { year: 'numeric', month: 'numeric', day: 'numeric', hour: 'numeric', minute: 'numeric' };
                } else {
                    options = { year: 'numeric', month: 'numeric', day: 'numeric' };
                }

                return val.toLocaleDateString('en-US', options)
            }
        },
        computed: {
            limitData () {
                let data = [...this.list];

                const sortColumns = this.sortColumns;
                if (this.sortType === 'asc') {
                    data = data.sort((a, b) => {
                        return a[sortColumns] > b[sortColumns] ? 1 : -1;
                    });
                }
                if (this.sortType === 'desc') {
                    data = data.sort((a, b) => {
                        return a[sortColumns] < b[sortColumns] ? 1 : -1;
                    });
                }

                if (this.filterType && this.filterType.length) {
                    data = data.filter(item => {
                        return this.filterType.indexOf(item.status) >= 0;
                    });
                }

                const selectedNames = this.selectedData.map(item => item.name);
                data.map(item => {
                    item._checked = selectedNames.indexOf(item.name) >= 0;
                    return item;
                });

                return data;
            },
            dataWithPage () {
                const data = this.limitData;
                const start = this.current * this.size - this.size;
                const end = start + this.size;
                return [...data].slice(start, end);
            },
            totalSelectedCount () {
                let count = 0;
                this.selectedData.forEach(item => {
                    count += item.count;
                });

                return count;
            },
            tableColumns () {
                const columns = [...this.columns];
                return columns.filter(item => item.show);
            }
        },
        methods: {
            getData () {
                if (this.loading) return;
                this.loading = true;
                const params = this.$parent.$parent.$refs.form.data;
                let query = {
                    status: params.status,
                    name: params.name,
                    phone: params.phone,
                    email: params.email,
                    date: params.date.length === 2 && params.date[0] !== '' && params.date[1] !== '' ? [this.$Date(params.date[0]).format('YYYY-MM-DD HH:mm'), this.$Date(params.date[1]).format('YYYY-MM-DD HH:mm')] : [],
                    page: this.current,
                    size: this.size,
                    sortType: this.sortType,
                    sortKey: this.sortColumns
                }
                GetUserLists(query).then(res => {
                    console.log(res.content)
                    this.handleClearSelect();
                    this.list = res.content.lists;
                    this.total = res.content.total;
                    this.untouched = res.content.untouched
                    this.loading = false;
                });
            },

            handleSortChange ({ key, order }) {
                // 将排序保存到数据
                this.sortColumns = key;
                this.sortType = order;
                this.current = 1;
            },
            handleFilterChange () {
                this.current = 1;
            },
            handleSelect (selection, row) {
                this.selectedData.push(row);
            },
            handleSelectCancel (selection, row) {
                const index = this.selectedData.findIndex(item => item.name === row.name);
                this.selectedData.splice(index, 1);
            },
            handleSelectAll (selection) {
                selection.forEach(item => {
                    if (this.selectedData.findIndex(i => i.name === item.name) < 0) {
                        this.selectedData.push(item);
                    }
                });
            },
            handleSelectAllCancel () {
                const selection = this.dataWithPage;
                selection.forEach(item => {
                    const index = this.selectedData.findIndex(i => i.name === item.name);
                    if (index >= 0) {
                        this.selectedData.splice(index, 1);
                    }
                });
            },
            handleClearSelect () {
                this.selectedData = [];
            },
            handleClickItem (name) {
                if (name === 'delete') {
                    this.selectedData.forEach(item => {
                        const index = this.list.findIndex(i => i.name === item.name);
                        if (index >= 0) {
                            this.list.splice(index, 1);
                        }
                    });
                    this.selectedData = [];
                }
            },
            handleOpenCreate () {
                // this.updateIndex = -1;
                // this.createData.desc = '';
                // this.showCreate = true;
                this.$Modal.info({
                    title: 'In Developing',
                    content: 'Create method is in developing'
                });
            },
            // 新增数据
            handleCreate () {

            },
            handleGetSlip (id) {
                console.log(this.list)
                return this.list.find(item => item.ID === id);
            },
            handleUpdate (id) {
                if (this.loading) return;
                this.$Modal.confirm({
                    title: 'Contacted Confirm',
                    content: `Have You Already Contacted To Patient <strong>${this.handleGetSlip(id).Name}</strong> Yet?`,
                    loading: true,
                    okText: 'Yes!',
                    onOk: () => {
                        UpdateUserStatus(id).then(res => {
                            this.$Modal.remove();
                            if (res.status === 200) {
                                this.$Message.success(res.msg);
                                this.getData();
                            } else {
                                this.$Message.error(res.msg);
                            }
                        });
                    }
                });
            },
            handleChangeTableSize (size) {
                this.tableSize = size;
            },
            handleFullscreen () {
                this.tableFullscreen = !this.tableFullscreen;
                this.$emit('on-fullscreen', this.tableFullscreen);
            },
            handleRefresh () {
                this.getData();
            },
            handleResetColumn () {
                this.columns = this.columns.map(item => {
                    const newItem = item;
                    newItem.show = true;
                    return newItem;
                });
            },
            handleChangePageSize (size) {
                this.size = size;
                this.getData();
            },
            handleChangePage (page) {
                this.current = page;
                this.getData();
            },
            handleClickView (id) {
                let slip = this.handleGetSlip(id)
                this.$Modal.info({
                    title: slip.Name + ' ' + 'Photo',
                    content: `<img src="${slip.Photo}" style="width: 100%"/>`,
                    className: 'aaaaaa'
                });
            }
        }
    }
</script>
