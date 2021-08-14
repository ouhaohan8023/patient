<template>
    <div class="page-account">
        <div class="page-account-container page-account-container-result">
            <div class="page-account-top">
                <div class="page-account-top-logo">
                    <img src="@/assets/images/logo.png" alt="logo">
                </div>
                <div class="page-account-top-desc">Done. Patient Registration</div>
            </div>
            <Result type="success" :title="title" :desc="desc">
                <div slot="actions">
                    <Button size="large" to="https://www.donefirst.com/">Go To Official Website</Button>
                </div>
            </Result>
        </div>
        <i-copyright />
    </div>
</template>
<script>
    import iCopyright from '@/components/copyright';

    export default {
        components: { iCopyright },
        data () {
            return {
                title: 'Register Successfully',
                desc: 'Thanks for your time, Our staff will contact you as soon as possible',
                time: 0
            }
        },
        beforeRouteEnter (to, from, next) {
            next(vm => {
                vm.prevRoute = from
            })
        },
        mounted () {
            if (this.prevRoute.name !== 'register') {
                this.$router.push({ name: 'register' })
            }
        }
    }
</script>
