<script lang=typescript>    
    import type { InstanceData, Day } from "../../openapi";

    import { Service } from '../../openapi';

    import Modal from "../Modal.svelte";
    import Button from "../model-input/Button.svelte";
    import AssistantInput from "./AssistantInput.svelte";
    import DatePicker from "./DatePicker.svelte";
    import HolidayPicker from "./HolidayPicker.svelte";

    export let data: InstanceData

    interface CustomDate {
        day: number,
        month: number,
        year: number
    }

    function handleSubmit() {
        Service.postInstanceData(data)
    }

    function addDays(date: CustomDate, days): CustomDate {
        var result: Date = new Date(date.year, date.month-1, date.day);
        result.setDate(result.getDate() + days);
        return {day: result.getDate(), month: result.getMonth()+1, year: result.getFullYear()};
    }

    function removeWeek() {
        data.days = data.days.slice(0, data.days.length - 7)
        handleSubmit()
    }

    function addWeek() {
        console.log(addDays(data.days[data.days.length-1].date, 1))
        let newDays = new Array(7)
            .fill(0)
            .map((_,i) => {
                return {
                    id: data.days.length + i + 1,
                    date: addDays(data.days[data.days.length-1].date, i+1),
                    is_holiday: false,
                }
            })
        data.days = data.days.concat(newDays)
        handleSubmit()

    }

    function setStartDate(startDate: Date) {
        let newDays = data.days.map((_,i) => {
            let date = addDays({day: startDate.getDate(), month: startDate.getMonth() + 1, year: startDate.getFullYear()}, i)
            return {
                id: i+1,
                date,
                is_holiday: data.days.some(d => d.date.day === date.day && d.date.month == date.month && d.date.year === date.year) && data.days.find(d => d.date.day === date.day && d.date.month == date.month && d.date.year === date.year).is_holiday
            }
        }) 
        data.days = newDays
    }

</script>

<main>
    <div class="flex flex-col space-y-2"> 
        <p class="font-semibold text-sm cursor-default"> Instance data </p>
        <p class="mt-4 font-semibold text-xs text-gray-500 cursor-default"> Start date </p>
        <DatePicker bind:currentDate={data.days[0].date} {handleSubmit} {setStartDate}/>
        <p class="mt-4 font-semibold text-xs text-gray-500 cursor-default"> Number of weeks </p>
        <div class="flex flex-row justify-start space-x-2 items-center">
            {#if data.days.length > 7}
            <Button callback={() => removeWeek()}> - </Button>
            {:else}
            <button disabled class="bg-gray-500 text-white font-bold px-2 py-1 rounded text-xs border-transparent">
                -
            </button>
            {/if}
            <p class="font-semibold text-sm text-black-500 cursor-default"> {data.days.length / 7} </p>
            <Button callback={() => addWeek()}> + </Button>
        </div>
        <p class="mt-4 font-semibold text-xs text-gray-500 cursor-default"> End date: </p>
        <p class="mt-4 font-semibold text-xs text-gray-500 cursor-default">
            {new Date(data.days[data.days.length-1].date.year, data.days[data.days.length-1].date.month-1, data.days[data.days.length-1].date.day).toDateString()}
        </p>
        <HolidayPicker bind:days={data.days} {handleSubmit}/>
        <Modal>
            <div class="mt-4" slot="trigger" let:open>
                <Button callback={open}> Edit assistants </Button>
            </div>
            <div class="my-2" slot="header">
                <h1 class="font-bold"> Edit assistants </h1>
            </div>
            <div class="my-2" slot="content">
                <AssistantInput bind:data={data}/>
            </div>
            <div class="my-5 flex flex-row justify-end space-x-2" slot="footer" let:close>
                <Button primary={false} callback={close}> Close </Button>
                <Button callback={() => {handleSubmit(); close()}}> Submit </Button>
            </div>
        </Modal>
    </div>
</main>