<template>
    <div v-resize="handleResize">
        <base-info ref="baseInfo" />
        <Card :bordered="false" dis-hover class="dashboard-console-visit">
            <div slot="title">
                <Row type="flex" justify="center" align="middle">
                    <Col span="8">
                        <Avatar icon="ios-podium" size="small" v-color="'#1890ff'" v-bg-color="'#e6f7ff'" />
                        <span class="ivu-pl-8">访问量</span>
                    </Col>
                    <Col span="16" class="ivu-text-right">
                        <RadioGroup v-model="visitType" type="button" class="ivu-mr-8" @on-change="handleChangeVisitType">
                            <Radio label="day">今日</Radio>
                            <Radio label="month">本月</Radio>
                            <Radio label="year">全年</Radio>
                        </RadioGroup>
                        <DatePicker v-model="visitDate" type="daterange" placement="bottom-end" placeholder="Select date" style="width: 200px"></DatePicker>
                    </Col>
                </Row>
            </div>
            <div>
                <visit-chart ref="visitChart" />
            </div>
        </Card>
    </div>
</template>
<script>
    import baseInfo from './base-info';
    import gridMenu from './grid-menu';
    import visitChart from './visit-chart';
    import hotSearch from './hot-search';
    import userPreference from './user-preference';

    export default {
        name: 'dashboard-console',
        components: { baseInfo, gridMenu, visitChart, hotSearch, userPreference },
        data () {
            return {
                visitType: 'day', // day, month, year
                visitDate: [(new Date()), (new Date())]
            }
        },
        methods: {
            handleChangeVisitType (val) {
                if (val === 'day') {
                    this.visitDate = [(new Date()), (new Date())];
                } else if (val === 'month') {
                    this.visitDate = [(new Date() - 86400000 * 30), (new Date())];
                } else if (val === 'year') {
                    this.visitDate = [(new Date() - 86400000 * 365), (new Date())];
                }
            },
            // 监听页面宽度变化，刷新表格
            handleResize () {
                this.$refs.baseInfo.handleResize();
                this.$refs.visitChart.handleResize();
                this.$refs.hotSearch.handleResize();
            }
        }
    }
</script>
<style lang="less">
    .dashboard-console-visit{
        .ivu-radio-group-button .ivu-radio-wrapper{
            border: none !important;
            box-shadow: none !important;
            padding: 0 12px;
        }
        .ivu-radio-group-button .ivu-radio-wrapper:before, .ivu-radio-group-button .ivu-radio-wrapper:after{
            display: none;
        }
    }
</style>
