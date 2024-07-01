<template>
	<router-view />
</template>

<script>
</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
}

.grp-container {
  width: calc(100% - 40px);
  height: calc(100vh - 40px);
  padding: 20px;
}

button{
  min-width: 100px;
  letter-spacing: 1px;
  height: 30px;
}

.button{
  border: 0;
  color: var(--color_text_default);
  background-color: #e9f0f0;
  font-weight: bold;
  font-size: 14px;

  &:hover {
    background-color: #132F51;
    color: #FFFFFF;
  }
}
</style>
