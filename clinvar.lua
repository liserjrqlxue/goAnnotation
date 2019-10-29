--
-- Created by IntelliJ IDEA.
-- User: wangyaoshen
-- Date: 2019/10/29
-- Time: 15:32
-- To change this template use File | Settings | File Templates.
--


CLNREVSTAT = {}
CLNREVSTAT['practice_guideline']='four'
CLNREVSTAT['reviewed_by_expert_panel']='three'
CLNREVSTAT['criteria_provided,_multiple_submitters,_no_conflicts']='two'
CLNREVSTAT['criteria_provided,_conflicting_interpretations']='one'
CLNREVSTAT['criteria_provided,_single_submitter']='one'
CLNREVSTAT['no_assertion_for_the_individual_variant']='none'
CLNREVSTAT['no_assertion_criteria_provided']='none'
CLNREVSTAT['no_assertion_provided']='none'

function countGoldStars(reviewStatus)
    return CLNREVSTAT[reviewStatus]
end
