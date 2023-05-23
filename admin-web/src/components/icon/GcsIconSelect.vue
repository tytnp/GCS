<template>
    <n-select @update:value="handleUpdateValue"
              v-model:value="iconName" placeholder="icon"
              :options="options" :render-label="renderLabel" filterable>
    </n-select>
</template>

<script setup>
import {iconMap} from './icons.js'

function handleUpdateValue (value){
    emit('update:value', value)
}

const emit = defineEmits(['update:value'])
const props = defineProps(
    {
        value: {
            type: String,
        }
    }
)

watch(()=>props.value,(newVal,oldValue)=>{
  iconName.value=newVal
})

const iconName = ref(props.value)

const renderLabel = (option) => {
    return [
        h(
            NIcon,
            {
                style: {
                    verticalAlign: "-0.15em",
                    marginRight: "4px"
                }
            },
            {
                 default: () => h(iconMap[option.label])
            }
        ),
        option.label
    ]
}

const options = []

for (const iconName in iconMap) {
    options.push(
        {
            label: iconName,
            value: iconName
        }
    );
}

</script>
