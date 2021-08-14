<template>
    <Form ref="form" :model="data" :rules="rules" :label-width="labelWidth" :label-position="labelPosition">
        <Row :gutter="24" type="flex" justify="end">
            <Col v-bind="grid">
                <FormItem label="Name：" prop="name" label-for="name">
                    <Input v-model="data.name" placeholder="Name" element-id="name" />
                </FormItem>
            </Col>
            <Col v-bind="grid">
              <FormItem label="Phone：" prop="phone" label-for="phone">
                <Input v-model="data.phone" placeholder="Phone" element-id="phone" />
              </FormItem>
            </Col>
          <Col v-bind="grid">
            <FormItem label="Email：" prop="email" label-for="email">
              <Input v-model="data.email" placeholder="Email" element-id="email" />
            </FormItem>
          </Col>
            <Col v-bind="grid">
                <FormItem label="Status：" prop="status" label-for="status">
                    <Select v-model="data.status" placeholder="Status" element-id="status">
                        <Option :value="0">Waiting</Option>
                        <Option :value="1">Contacted</Option>
                    </Select>
                </FormItem>
            </Col>
            <Col v-bind="grid" class="ivu-text-right">
                <FormItem>
                    <Button type="primary" @click="handleSubmit">Search</Button>
                    <Button class="ivu-ml-8" @click="handleReset">Reset</Button>
                </FormItem>
            </Col>
        </Row>
    </Form>
</template>
<script>
    import { mapState } from 'vuex';

    export default {
        data () {
            return {
                grid: {
                    xl: 6,
                    lg: 6,
                    md: 12,
                    sm: 24,
                    xs: 24
                },
                collapse: false,
                data: {
                    name: '',
                    status: '',
                    count: null,
                    date: '',
                    phone: '',
                    email: ''
                },
                rules: {

                }
            }
        },
        computed: {
            ...mapState('admin/layout', [
                'isMobile'
            ]),
            labelWidth () {
                return this.isMobile ? undefined : 100;
            },
            labelPosition () {
                return this.isMobile ? 'top' : 'right';
            }
        },
        methods: {
            handleSubmit () {
                this.$emit('on-submit', this.data);
            },
            handleReset () {
                this.$refs.form.resetFields();
                this.$emit('on-reset');
            }
        }
    }
</script>
