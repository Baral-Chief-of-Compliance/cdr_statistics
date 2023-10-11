<template>
    <VuePlotly 
        :data="[{
            labels : labels,
            values : values,
            type: 'pie'
        }]"

        :layout="{
            title: title
        }"
    >
    </VuePlotly>
</template>

<script>
import { VuePlotly} from 'vue3-plotly';


export default{
    components: {
        VuePlotly
    },

    props: {
        labels: Array,
        values: Array,
        title: String
    },
}
</script>