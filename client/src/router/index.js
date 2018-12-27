import Vue from 'vue'
import Router from 'vue-router'
import Targeted from '@/components/Targeted'
import Administrates from '@/components/Administrates'
import Responses from '@/components/Responses'
import Explorer from '@/components/Explorer'
import QuestionnaireDetails from '@/components/QuestionnaireDetails'
import Results from '@/components/Results'
import EditResponse from '@/components/EditResponse'
import NotFound from '@/components/Utils/NotFound'

Vue.use(Router)

export default new Router({
  mode: 'history',
  props: {
    traqId: String
  },
  routes: [
    {
      path: '/',
      redirect: '/targeted'
    },
    {
      path: '/targeted',
      name: 'Targeted',
      component: Targeted,
      props: { traqId: String(this.traqId) }
    },
    {
      path: '/administrates',
      name: 'Administrates',
      component: Administrates
    },
    {
      path: '/responses',
      name: 'Responses',
      component: Responses
    },
    {
      path: '/explorer',
      name: 'Explorer',
      component: Explorer
    },
    {
      path: '/questionnaires/:id',
      name: 'QuestionnaireDetails',
      component: QuestionnaireDetails
    },
    {
      path: '/results/:id',
      name: 'Results',
      component: Results
    },
    {
      path: '/questionnaires/:id/new-response',
      name: 'EditResponse',
      component: EditResponse
    },
    {
      path: '*',
      name: 'NotFound',
      component: NotFound
    }
  ]
})
