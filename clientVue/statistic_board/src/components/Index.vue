<template>
    <v-container class="mt-2">


        <v-row>

            <v-col>
                <v-row>
                    <v-col>
                        <v-card class="pb-2">
                            <v-tabs
                                v-model="tab"
                                bg-color="primary"

                            >
                                <v-row>
                                    <v-col>
                                        <v-tab @click="this.visible_charts = false; this.visible_charts_interval = false" value="1">сегодня</v-tab>
                                    </v-col>
                                    <v-col>
                                        <v-tab @click="this.visible_charts = false; this.visible_charts_interval = false" value="2">неделя</v-tab>
                                    </v-col>
                                    <v-col>
                                        <v-tab @click="this.visible_charts = false; this.visible_charts_interval = false" value="3">месяц</v-tab>
                                    </v-col>
                                    <v-col>
                                        <v-tab value="4">произвольный промежуток</v-tab>
                                    </v-col>
                                </v-row>      
                            </v-tabs>

                            <v-card-text>
                                <v-window v-model="tab">
                                    <v-window-item value="1">
                                        <div v-if="this.select_department == null" class="text-h5">Выберите отдел</div>
                                        <v-btn @click="this.getStatisticsToday()" v-if="this.select_department != null" class="mt-5" color="primary" block>получить статистику за сегодня</v-btn>

                                    </v-window-item>

                                    <v-window-item value="2">
                                        <div v-if="this.select_department == null" class="text-h5">Выберите отдел</div>
                                        <v-btn @click="this.getStatisticsWeek()" v-if="this.select_department != null" class="mt-5" color="primary" block>получить статистику за неделю</v-btn>
                                    </v-window-item>

                                    <v-window-item value="3">
                                        <div v-if="this.select_department == null" class="text-h5">Выберите отдел</div>
                                        <v-btn @click="this.getStatisticsMonth()" v-if="this.select_department != null" class="mt-5" color="primary" block>получить статистику за месяц</v-btn>
                                    </v-window-item>

                                    <v-window-item value="4">
                                        <div v-if="this.select_department == null" class="text-h5">Выберите отдел</div>

                                        <v-row v-if="this.select_department != null">
                                            <v-col>
                                                <v-row>
                                                    <v-col>
                                                        <div class="text-h6">Дата</div>
                                                        <v-date-picker
                                                            v-model="date_start"
                                                            color="primary"
                                                        >
                                                        </v-date-picker>
                                                    </v-col>
                                                    <v-col>
                                                        <div class="text-h6">Время</div>
                                                        <input v-model="time_start" class="time_input" type="time" />                                      
                                                    </v-col>
                                                </v-row>
                                            </v-col>

                                            <v-col>
                                                <v-row>
                                                    <v-col>
                                                        <div class="text-h6">Дата</div>
                                                        <v-date-picker
                                                            v-model="date_end"
                                                            color="primary"
                                                        >
                                                        </v-date-picker>
                                                    </v-col>
                                                    <v-col>
                                                        <div class="text-h6">Время</div>
                                                        <input v-model="time_end" class="time_input" type="time" />
                                                    </v-col>
                                                </v-row>
                                            </v-col>
                                        </v-row>

                                        <v-btn 
                                            class="mt-5" 
                                            v-if="this.date_start != '' && this.date_end != '' && this.time_end != '' && this.time_start != '' && this.select_department != null" block color="primary"
                                            @click="this.getStatistics()"
                                        >
                                            Получить статистику за произвольный промежуток
                                        </v-btn>
                                    </v-window-item>
                                </v-window>  
                            </v-card-text>
                        </v-card>
                    </v-col>

                </v-row>
            </v-col>




            <!-- chose departmnet and count statistics -->
            <v-col >
                <v-row>
                    <v-card width="700" class="mt-3">
                        <v-card-title>Отдел</v-card-title>
                        <v-card-text>
                            <v-select
                                label="Название отдела"
                                :items="departments_labels"
                                v-model="select_department"
                                variant="solo"
                            >
                            </v-select>
                        </v-card-text>

                    </v-card>
                </v-row>

                <v-row v-if="visible_charts_interval && tab == 4">
                    <v-col>
                        <PieChart 
                            :labels="['Отвечено', 'Не отвечено']"
                            :values="[this.inbounds_answered, this.inbounds_no_answered]"
                            title="Внутренние"
                        />
                    </v-col>
                    <v-col>
                        <PieChart 
                            :labels="['Отвечено', 'Не отвечено']"
                            :values="[this.outbounds_answered, this.outbounds_no_answered]"
                            title="Внешние"
                        />
                    </v-col>

                </v-row>

            </v-col>

        </v-row>

        <v-row v-if="visible_charts && tab != 4">
            <v-col>
                <PieChart 
                    :labels="['Отвечено', 'Не отвечено']"
                    :values="[this.inbounds_answered, this.inbounds_no_answered]"
                    title="Внутренние"
                />
            </v-col>
            <v-col>
                <PieChart 
                    :labels="['Отвечено', 'Не отвечено']"
                    :values="[this.outbounds_answered, this.outbounds_no_answered]"
                    title="Внешние"
                />
            </v-col>
        </v-row>

    </v-container>

</template>

<script>
import axios from 'axios';
import PieChart from './PieChart.vue';


export default{
    components: {
        PieChart
    },
    data: () => ({
        
           date_start: "",
           time_start: "",
           date_end: "",
           time_end: "",
           select_department: null,
           start_result: "",
           tab: null,
           departments_labels: [
            "Отдел организации закупок",
            "Отдел правового и кадрового обеспечения",
            "Отдел рынка труда, статистики и взаимодействия с работодателями",
            "Отдел бухгалтерского учета",
            "Отдел социальных выплат и экономического анализа",
            "Отдел методологического сопровождения, контроля оказания услуг, улучшения качества клиентского опыта",
            "Отдел оказания услуг по сопровождению и подбору персонала",
            "Пресс-служба",
            "Административно-хозяйственный отдел",
            "Отдел информационных систем и автоматизации процессов"
           ],

           departments_phone: {
            "Отдел организации закупок": "770,771,772,773",
            "Отдел правового и кадрового обеспечения": "216,710,711,712,713,714,715,717",
            "Отдел рынка труда, статистики и взаимодействия с работодателями": "775,751,752,753,754,755,756",
            "Отдел бухгалтерского учета": "720,721,722,724,725,723,726,727",
            "Отдел социальных выплат и экономического анализа": "731,734,733,735,732",
            "Отдел методологического сопровождения, контроля оказания услуг, улучшения качества клиентского опыта": "780,781,783,782",
            "Отдел оказания услуг по сопровождению и подбору персонала": "760,761,762,763,764,766,767,768,728",
            "Пресс-служба": "790,791",
            "Административно-хозяйственный отдел": "740,746,748,747,742,744,745,741,743",
            "Отдел информационных систем и автоматизации процессов": "700,701,702,703,704,705"
           },

           visible_charts: false,
           visible_charts_interval: false,

           inbounds_answered: 0,
           inbounds_no_answered: 0,
           outbounds_answered: 0,
           outbounds_no_answered: 0
        
    }),

    methods: {
       FormatTimeDate(strDate, strTime){
        const date = new Date(strDate)
        const [month, day, year] = [
            date.getMonth() + 1,
            date.getDate(),
            date.getFullYear()
        ]

        return `${year}-${month}-${day} ${strTime}:00`
       },

       getStatistics(){
            const starttime = this.FormatTimeDate(this.date_start, this.time_start);
            const endtime = this.FormatTimeDate(this.date_end, this.time_end);
            const departments_map = new Map(Object.entries(this.departments_phone));
            const numbers = departments_map.get(this.select_department);

            console.log(numbers);
            console.log(starttime);
            console.log(endtime);

            axios.post('http://localhost:8080/statistics', {
            numbers:  numbers,
            starttime: starttime,
            endtime: endtime
            }).then(response => (
                this.inbounds_answered = response.data.inbounds_answered,
                this.inbounds_no_answered = response.data.inbounds - this.inbounds_answered,
                this.outbounds_answered = response.data.outbounds_answered,
                this.outbounds_no_answered = response.data.outbounds - this.outbounds_answered,
                this.visible_charts_interval = true
            ))
       },

       getStatisticsToday(){
            const today = new Date()
            const [month, day, year] = [
                today.getMonth() + 1,
                today.getDate(),
                today.getFullYear()
            ]

            const starttime = `${year}-${month}-${day} 00:00:00`;
            const endtime = `${year}-${month}-${day} 23:59:00`;
            const departments_map = new Map(Object.entries(this.departments_phone));
            const numbers = departments_map.get(this.select_department);

            axios.post('http://localhost:8080/statistics', {
                numbers:  numbers,
                starttime: starttime,
                endtime: endtime
            }).then(response => (
                this.inbounds_answered = response.data.inbounds_answered,
                this.inbounds_no_answered = response.data.inbounds - this.inbounds_answered,
                this.outbounds_answered = response.data.outbounds_answered,
                this.outbounds_no_answered = response.data.outbounds - this.outbounds_answered,
                this.visible_charts = true
            ));
       },

       getStatisticsWeek(){
        const today = new Date()
        const dayWeekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000);
        const departments_map = new Map(Object.entries(this.departments_phone));
        const numbers = departments_map.get(this.select_department);

        const starttime = this.FormatTimeDate(dayWeekAgo, '00:00');
        const endtime = this.FormatTimeDate(today, '23:59');

        console.log(starttime)
        console.log(endtime)

        axios.post('http://localhost:8080/statistics', {
                numbers:  numbers,
                starttime: starttime,
                endtime: endtime
            }).then(response => (
                this.inbounds_answered = response.data.inbounds_answered,
                this.inbounds_no_answered = response.data.inbounds - this.inbounds_answered,
                this.outbounds_answered = response.data.outbounds_answered,
                this.outbounds_no_answered = response.data.outbounds - this.outbounds_answered,
                this.visible_charts = true
            )) 
       },

       getStatisticsMonth(){
        const today = new Date();

        const [todayMonth, todayDay, todayYear] = [
                today.getMonth() + 1,
                today.getDate(),
                today.getFullYear()
            ]


        const dayMonthAgo = today.setMonth(today.getMonth() - 1);

        const starttime = this.FormatTimeDate(dayMonthAgo, '00:00');
        const endtime = `${todayYear}-${todayMonth}-${todayDay} 23:59:00`;

        console.log(starttime)
        console.log(endtime)

        const departments_map = new Map(Object.entries(this.departments_phone));
        const numbers = departments_map.get(this.select_department);

        axios.post('http://localhost:8080/statistics', {
                numbers:  numbers,
                starttime: starttime,
                endtime: endtime
            }).then(response => (
                this.inbounds_answered = response.data.inbounds_answered,
                this.inbounds_no_answered = response.data.inbounds - this.inbounds_answered,
                this.outbounds_answered = response.data.outbounds_answered,
                this.outbounds_no_answered = response.data.outbounds - this.outbounds_answered,
                this.visible_charts = true
            )) 
        
       }
    }
}
</script>

<style scoped>


</style>