<script lang=typescript>

    import type { Schedule, InstanceData, Assistant, IndividualSchedule } from '../../openapi';
    import { AssistantType } from '../../openapi';
    
    import AssistantHeader from './AssistantHeader.svelte'
    import DayHeader from './DayHeader.svelte';
    import AssignmentBox from './AssignmentBox.svelte';

    import { onMount } from 'svelte';

    export let schedule: Schedule;
    export let data: InstanceData;

    let max_workload = Math.max(...schedule.individual_schedules.map((a) => a.workload))
    let min_workload = Math.min(...schedule.individual_schedules.map((a) => a.workload))

    function getAssistant(id: number): Assistant {
        return data.assistants.find(a => a.id === id);
    }

    function getIs(id: number): IndividualSchedule {
        return schedule.individual_schedules.find(s => s.assistant_id === id);
    }

    const types: AssistantType[] = Object.values(AssistantType)

    onMount(async () => {
		let isSyncingLeftScroll: boolean = false;
        let isSyncingTopScroll: boolean = false;
        let isSyncingCenterScroll: boolean = false;
        let leftDiv: HTMLElement = document.getElementById('assistantlist');
        let topDiv: HTMLElement = document.getElementById('dayheaders');
        let centerDiv: HTMLElement = document.getElementById('schedule');

        leftDiv.onscroll = function() {
            if (!isSyncingLeftScroll) {
                isSyncingCenterScroll = true;
                centerDiv.scrollTop = leftDiv.scrollTop;
            }
            isSyncingLeftScroll = false;
        }

        topDiv.onscroll = function() {
            if (!isSyncingTopScroll) {
                isSyncingCenterScroll = true;
                centerDiv.scrollLeft = topDiv.scrollLeft;
            }
            isSyncingTopScroll = false;
        }

        centerDiv.onscroll = function() {
            if (!isSyncingCenterScroll) {
                isSyncingLeftScroll = true;
                isSyncingTopScroll = true;
                leftDiv.scrollTop = centerDiv.scrollTop;
                topDiv.scrollLeft = centerDiv.scrollLeft;
            }
            isSyncingCenterScroll = false;

        }
	});

</script>

<div class="flex flex-row space-x-10 justify-end w-full font-bold text-sm">
    <p> fairness score: {schedule.fairness_score.toFixed(2)} </p>
    <p> JAEV fairness score: {schedule.jaev_fairness_score.toFixed(2)} </p>
</div>

<div class="flex flex-row flex-nowrap space-x-2 w-full overflow-y-scroll">
    <div class="flex flex-col h-full">
        <div class="flex flex-none h-10"/> <!--placeholder-->
        <div id="assistantlist" class="scrollbar-hidden flex flex-row h-full overflow-y-scroll">
            <div  class="flex flex-none flex-col h-full space-y-1">
                {#each types as type}
                    {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
                        <AssistantHeader {assistant} workload={getIs(assistant.id).workload} {max_workload} {min_workload}/>
                    {/each}
                {/each}
            </div>
        </div>
    </div>
    <div class="flex flex-col h-full overflow-scroll space-y-2">
        <div id="dayheaders" class="scrollbar-hidden flex flex-none flex-row overflow-x-scroll space-x-1">
            {#each data.days as day}
                <DayHeader {day}/>
                {#if day.id % 7 === 0}
                    <div class="flex flex-none w-4"/>
                {/if}
            {/each}
        </div>
        <div id="schedule" class="scrollable flex flex-col h-full w-full overflow-x-scroll overflow-y-scroll space-y-1">
            {#each types as type}
                {#each data.assistants.filter((a) => a.type == type) as assistant (assistant.id)}
                    <div class="flex flex-row w-full space-x-1">
                        {#each data.days as day}
                            <AssignmentBox assignment={getIs(assistant.id).assignments.find(a => a.day_nb === day.id)} free_day={assistant.free_days.includes(day.id)}/>
                             {#if day.id % 7 === 0}
                                <div class="flex flex-none w-4"/>
                            {/if}
                        {/each}
                    </div>
                {/each}
            {/each}
        </div>  
    </div>
</div>