<template>
    <v-container>
        <v-row>

            <v-col>
                <v-card>
                    <v-card-title>Начало промежутка</v-card-title>
                    <v-card-text>
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
                    </v-card-text>
                </v-card>
            </v-col>

            

            <v-col>
                <v-card>
                    <v-card-title>Конец промежутка</v-card-title>
                    <v-card-text>
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
                    </v-card-text>
                </v-card>
            </v-col>

            <v-col>
                <v-row>
                    <v-card width="700">
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

                <v-row v-if="visible_charts">
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

        <v-row>
            <v-btn 
                block
                color="primary"
                @click="this.getStatistics()"
            >
                Получить статистику
            </v-btn>
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
           select_department: "",
           start_result: "",
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

           inbounds_answered: 0,
           inbounds_no_answered: 0,
           outbounds_answered: 0,
           outbounds_no_answered: 0
        
    }),

    methods: {
       FormatTimeDate(strDate, strTime){
        const date = new Date(strDate)
        const [month, day, year] = [
            date.getMonth(),
            date.getDate(),
            date.getFullYear()
        ]

        return `${year}-${month}-${day} ${strTime}:00`
       },

       getStatistics(){
            const starttime = this.FormatTimeDate(this.date_start, this.time_start);
            const endtime = this.FormatTimeDate(this.date_end, this.time_end);
            const departments_map = new Map(Object.entries(this.departments_phone))
            const numbers = departments_map.get(this.select_department)

            console.log(numbers)
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

            

       }
    }
}
</script>

<style scoped>


</style>