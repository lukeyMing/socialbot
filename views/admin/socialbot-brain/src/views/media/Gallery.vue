<template>
    <div class="gallery">
        <masonry :cols="{default: 4, 1200: 4, 768: 2, 576: 1}" :gutter="{default: '10px', 1200: '10px', 768: '10px'}">
            <gallery-card v-for="g in gallery" :key="g.Id" :category-list="category" :media="g"></gallery-card>
        </masonry>
    </div>
</template>

<script>


    import Vue from 'vue'
    import VueMasonry from 'vue-masonry-css'
    Vue.use(VueMasonry);
    import {listGalleryApi} from "../../api/gallery"
    import {listCategoryWithTagsApi} from "../../api/category";
    import GalleryCard from "../../components/media/GalleryCard"

    export default {
        name: "Gallery",
        created() {
            this.listCategory();
            this.listGallery();
        },
        data(){
            return {
                category:[],
                gallery:[],
            }
        },
        components:{
            GalleryCard
        },
        methods:{
            listCategory() {
                listCategoryWithTagsApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.category = res.data;
                })
            },
            listGallery() {
                listGalleryApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.gallery = res.data;
                })
            },
        },
    }
</script>

<style scoped>

</style>
