<template>
    <div class="recycle-scroller-item">
      <div class="mb-2xs card _card_e3hfs_5 _cardLink_86eyf_5" data-test-id="resources-list-item">
        <div class="_content_e3hfs_24">
          <div class="_header_e3hfs_16">
            <h2 class="n8n-heading _size-medium_1v2ig_28 _bold_1v2ig_5 _cardHeading_86eyf_13" data-test-id="workflow-card-name">
              {{ workflowName }}
            </h2>
          </div>
          <div class="_body_e3hfs_32">
            <div class="_cardDescription_86eyf_18">
              <span class="n8n-text _text-light_15q07_54 _size-small_15q07_28 _regular_15q07_9">
                <span>Last updated <span :title="updatedDateTooltip">{{ updatedDate }}</span> | </span>
                <span class="mr-2xs">{{ createdDate }}</span>
              </span>
            </div>
          </div>
        </div>

        <!--  STATUS -->
        <div>
          <div class="_cardActions_86eyf_24">
            <div></div>
            <div data-v-76d7fff1="" class="workflow-activator mr-s" data-test-id="workflow-card-activator">
              <div data-v-76d7fff1="" data-test-id="workflow-activator-status" class="_activeStatusText_1dwsh_5">
                <span data-v-76d7fff1="" class="n8n-text _success_15q07_66 _size-small_15q07_28 _bold_15q07_5">
                  {{ workflowStatus }}
                </span>
              </div>
              <!-- ... (rest of the code for switch and dropdown menu) ... -->
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'WorkflowCard',
    data() {
      return {
        workflowName: 'Test_Nouha CD to Clay',
        updatedDateTooltip: '4 October, 2023 @ 10:58 GMT',
        updatedDate: '3 days ago',
        createdDate: 'Created 15 August',
        workflowStatus: 'Active',
        // ... (any other necessary data properties) ...
      };
    },
    methods: {
      // ... (any necessary methods for actions like activating the workflow, opening the menu, etc.) ...
    },
  };
  </script>
  
  <style scoped>
._bold_15q07_5 {
    font-weight: var(--font-weight-bold)
}

._regular_15q07_9 {
    font-weight: var(--font-weight-regular)
}

._size-xlarge_15q07_13 {
    font-size: var(--font-size-xl);
    line-height: var(--font-line-height-xloose)
}

._size-large_15q07_18 {
    font-size: var(--font-size-m);
    line-height: var(--font-line-height-xloose)
}

._size-medium_15q07_23 {
    font-size: var(--font-size-s);
    line-height: var(--font-line-height-loose)
}

._size-small_15q07_28 {
    font-size: var(--font-size-2xs);
    line-height: var(--font-line-height-loose)
}

._size-xsmall_15q07_33 {
    font-size: var(--font-size-3xs);
    line-height: var(--font-line-height-compact)
}

._compact_15q07_38 {
    line-height: 1
}

._primary_15q07_42 {
    color: var(--color-primary)
}

._text-dark_15q07_46 {
    color: var(--color-text-dark)
}

._text-base_15q07_50 {
    color: var(--color-text-base)
}

._text-light_15q07_54 {
    color: var(--color-text-light)
}

._text-xlight_15q07_58 {
    color: var(--color-text-xlight)
}

._danger_15q07_62 {
    color: var(--color-danger)
}

._success_15q07_66 {
    color: var(--color-success)
}

._warning_15q07_70 {
    color: var(--color-warning)
}

._align-left_15q07_74 {
    text-align: left
}

._align-right_15q07_78 {
    text-align: right
}

._align-center_15q07_82 {
    text-align: center
}

._xlarge_yw9fy_5 {
    width: var(--font-size-xl) !important
}

._large_yw9fy_9 {
    width: var(--font-size-m) !important
}

._medium_yw9fy_13 {
    width: var(--font-size-s) !important
}

._small_yw9fy_17 {
    width: var(--font-size-2xs) !important
}

._xsmall_yw9fy_21 {
    width: var(--font-size-3xs) !important
}

.lds-ring {
    display: inline-block;
    position: relative;
    width: 48px;
    height: 48px
}

.lds-ring div {
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    display: block;
    position: absolute;
    width: 48px;
    height: 48px;
    border: 4px solid var(--color-foreground-xlight);
    border-radius: 50%;
    -webkit-animation: lds-ring 1.2s cubic-bezier(.5, 0, .5, 1) infinite;
    animation: lds-ring 1.2s cubic-bezier(.5, 0, .5, 1) infinite;
    border-color: var(--color-primary) transparent transparent transparent
}

.lds-ring div:nth-child(1) {
    -webkit-animation-delay: -.45s;
    animation-delay: -.45s
}

.lds-ring div:nth-child(2) {
    -webkit-animation-delay: -.3s;
    animation-delay: -.3s
}

.lds-ring div:nth-child(3) {
    -webkit-animation-delay: -.15s;
    animation-delay: -.15s
}

@-webkit-keyframes lds-ring {
    0% {
        -webkit-transform: rotate(0deg);
        transform: rotate(0)
    }

    to {
        -webkit-transform: rotate(360deg);
        transform: rotate(360deg)
    }
}

@keyframes lds-ring {
    0% {
        -webkit-transform: rotate(0deg);
        transform: rotate(0)
    }

    to {
        -webkit-transform: rotate(360deg);
        transform: rotate(360deg)
    }
}

._button_1hk7l_119 {
    display: inline-block;
    line-height: 1;
    white-space: nowrap;
    cursor: pointer;
    border: var(--border-width-base) var(--button-border-color, var(--color-primary)) var(--border-style-base);
    color: var(--button-color, var(--color-text-xlight));
    background-color: var(--button-background-color, var(--color-primary));
    font-weight: var(--font-weight-bold);
    border-radius: var(--button-border-radius, 4px);
    padding: var(--button-padding-vertical, var(--spacing-xs)) var(--button-padding-horizontal, var(--spacing-m));
    font-size: var(--button-font-size, var(--font-size-s));
    -webkit-appearance: none;
    text-align: center;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    outline: none;
    margin: 0;
    -webkit-transition: .3s;
    transition: .3s;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none
}

._button_1hk7l_119:hover {
    color: var(--button-hover-color, var(--color-text-xlight));
    border-color: var(--button-hover-border-color, var(--color-primary-shade-1));
    background-color: var(--button-hover-background-color, var(--color-primary-shade-1))
}

._button_1hk7l_119:focus {
    border-color: var(--button-focus-outline-color, var(--color-primary-tint-1));
    outline: 2px solid var(--button-focus-outline-color, var(--color-primary-tint-1))
}

._button_1hk7l_119:active,
._button_1hk7l_119._active_1hk7l_150 {
    color: var(--button-active-color, var(--color-text-xlight));
    border-color: var(--button-active-border-color, var(--color-primary-shade-1));
    background-color: var(--button-active-background-color, var(--color-primary-shade-1));
    outline: none
}

._button_1hk7l_119::-moz-focus-inner {
    border: 0
}

._button_1hk7l_119>i {
    display: none
}

._button_1hk7l_119>span {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._button_1hk7l_119 span+span {
    margin-left: var(--spacing-3xs)
}

._secondary_1hk7l_174 {
    --button-color: var(--color-primary);
    --button-border-color: var(--color-primary);
    --button-background-color: var(--color-background-xlight);
    --button-active-background-color: var(--color-primary-tint-2);
    --button-active-color: var(--color-primary);
    --button-active-border-color: var(--color-primary);
    --button-hover-background-color: var(--color-primary-tint-3);
    --button-hover-color: var(--color-primary);
    --button-hover-border-color: var(--color-primary);
    --button-focus-outline-color: var(--color-primary-tint-1)
}

._tertiary_1hk7l_187 {
    font-weight: var(--font-weight-bold) !important;
    --button-background-color: var(--color-background-xlight);
    --button-color: var(--color-text-dark);
    --button-border-color: var(--color-neutral-850);
    --button-active-background-color: var(--color-primary-tint-2);
    --button-active-color: var(--color-primary);
    --button-active-border-color: var(--color-primary);
    --button-hover-background-color: var(--color-neutral-950);
    --button-hover-color: var(--color-text-dark);
    --button-hover-border-color: var(--color-neutral-800);
    --button-focus-outline-color: hsla(var(--color-neutral-h), var(--color-neutral-s), var(--color-neutral-l), .2)
}

._success_1hk7l_206 {
    --button-background-color: var(--color-success);
    --button-color: var(--color-text-xlight);
    --button-border-color: var(--color-success);
    --button-active-background-color: var(--color-success-350);
    --button-active-border-color: var(--color-success-350);
    --button-hover-background-color: var(--color-success-450);
    --button-hover-border-color: var(--color-success-450);
    --button-focus-outline-color: hsla(var(--color-success-h), var(--color-success-s), var(--color-success-l), .33)
}

._warning_1hk7l_222 {
    --button-background-color: var(--color-warning);
    --button-color: var(--color-text-xlight);
    --button-border-color: var(--color-warning);
    --button-active-background-color: var(--color-warning-500);
    --button-active-border-color: var(--color-warning-500);
    --button-hover-background-color: var(--color-warning-650);
    --button-hover-border-color: var(--color-warning-650);
    --button-focus-outline-color: hsla(var(--color-warning-h), var(--color-warning-s), var(--color-warning-l), .33)
}

._danger_1hk7l_238 {
    --button-background-color: var(--color-danger);
    --button-color: var(--color-text-xlight);
    --button-border-color: var(--color-danger);
    --button-active-color: var(--color-text-xlight);
    --button-active-background-color: var(--color-danger-600);
    --button-active-border-color: var(--color-danger-600);
    --button-hover-background-color: var(--color-danger-700);
    --button-hover-border-color: var(--color-danger-700);
    --button-focus-outline-color: hsla(var(--color-danger-h), var(--color-danger-s), var(--color-danger-l), .33)
}

._xmini_1hk7l_258 {
    --button-padding-vertical: var(--spacing-4xs);
    --button-padding-horizontal: var(--spacing-3xs);
    --button-font-size: var(--font-size-3xs)
}

._xmini_1hk7l_258._square_1hk7l_263 {
    height: 22px;
    width: 22px
}

._mini_1hk7l_268 {
    --button-padding-vertical: var(--spacing-4xs);
    --button-padding-horizontal: var(--spacing-2xs);
    --button-font-size: var(--font-size-2xs)
}

._mini_1hk7l_268._square_1hk7l_263 {
    height: 22px;
    width: 22px
}

._small_1hk7l_278 {
    --button-padding-vertical: var(--spacing-3xs);
    --button-padding-horizontal: var(--spacing-xs);
    --button-font-size: var(--font-size-2xs)
}

._small_1hk7l_278._square_1hk7l_263 {
    height: 26px;
    width: 26px
}

._medium_1hk7l_288 {
    --button-padding-vertical: var(--spacing-2xs);
    --button-padding-horizontal: var(--spacing-xs);
    --button-font-size: var(--font-size-2xs)
}

._medium_1hk7l_288._square_1hk7l_263 {
    height: 30px;
    width: 30px
}

._large_1hk7l_298._square_1hk7l_263 {
    height: 42px;
    width: 42px
}

._xlarge_1hk7l_303 {
    --button-padding-vertical: var(--spacing-xs);
    --button-padding-horizontal: var(--spacing-s);
    --button-font-size: var(--font-size-m)
}

._xlarge_1hk7l_303._square_1hk7l_263 {
    height: 46px;
    width: 46px
}

._outline_1hk7l_316 {
    --button-color: var(--color-primary);
    --button-background-color: transparent;
    --button-disabled-background-color: transparent;
    --button-active-background-color: transparent
}

._outline_1hk7l_316._primary_1hk7l_322 {
    --button-color: var(--color-primary);
    --button-border-color: var(--color-primary);
    --button-active-background-color: var(--color-primary)
}

._outline_1hk7l_316._tertiary_1hk7l_187 {
    --button-color: var(--color-text-dark)
}

._outline_1hk7l_316._success_1hk7l_206 {
    --button-color: var(--color-success);
    --button-border-color: var(--color-success);
    --button-active-background-color: var(--color-success)
}

._outline_1hk7l_316._warning_1hk7l_222 {
    --button-color: var(--color-warning);
    --button-border-color: var(--color-warning);
    --button-active-background-color: var(--color-warning)
}

._outline_1hk7l_316._danger_1hk7l_238 {
    --button-color: var(--color-danger);
    --button-border-color: var(--color-danger);
    --button-active-background-color: var(--color-danger)
}

._text_1hk7l_346 {
    --button-color: var(--color-text-light);
    --button-border-color: transparent;
    --button-background-color: transparent;
    --button-active-color: var(--color-text-light);
    --button-active-background-color: transparent;
    --button-active-border-color: transparent;
    --button-hover-color: var(--color-text-light);
    --button-hover-background-color: transparent;
    --button-hover-border-color: transparent
}

._text_1hk7l_346._primary_1hk7l_322 {
    --button-color: var(--color-primary);
    --button-active-color: var(--color-primary);
    --button-hover-color: var(--color-primary)
}

._text_1hk7l_346._secondary_1hk7l_174 {
    --button-color: var(--color-primary-tint-1);
    --button-active-color: var(--color-primary-tint-1);
    --button-hover-color: var(--color-primary-tint-1)
}

._text_1hk7l_346._success_1hk7l_206 {
    --button-color: var(--color-success);
    --button-active-color: var(--color-success);
    --button-hover-color: var(--color-success)
}

._text_1hk7l_346._tertiary_1hk7l_187 {
    --button-hover-color: var(--color-primary)
}

._text_1hk7l_346._warning_1hk7l_222 {
    --button-color: var(--color-warning);
    --button-active-color: var(--color-warning);
    --button-hover-color: var(--color-warning)
}

._text_1hk7l_346._danger_1hk7l_238 {
    --button-color: var(--color-danger);
    --button-active-color: var(--color-danger);
    --button-hover-color: var(--color-danger)
}

._text_1hk7l_346:hover {
    text-decoration: underline
}

._loading_1hk7l_389 {
    position: relative;
    pointer-events: none
}

._loading_1hk7l_389:before {
    pointer-events: none;
    content: "";
    position: absolute;
    left: -1px;
    top: -1px;
    right: -1px;
    bottom: -1px;
    border-radius: inherit;
    background-color: #fff0
}

._disabled_1hk7l_405,
._disabled_1hk7l_405:hover,
._disabled_1hk7l_405:active,
._disabled_1hk7l_405:focus {
    cursor: not-allowed;
    background-image: none;
    color: var(--button-disabled-color, var(--color-text-dark));
    background-color: var(--button-disabled-background-color, var(--color-light-grey));
    border-color: var(--button-disabled-border-color, var(--color-light-grey))
}

._transparent_1hk7l_413 {
    --button-background-color: transparent;
    --button-active-background-color: transparent
}

._withIcon_1hk7l_418,
._icon_1hk7l_424 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._icon_1hk7l_424 svg {
    display: block
}

._block_1hk7l_433 {
    width: 100%
}

._float-left_1hk7l_437 {
    float: left
}

._float-right_1hk7l_441 {
    float: right
}

._bold_1v2ig_5 {
    font-weight: var(--font-weight-bold)
}

._regular_1v2ig_9 {
    font-weight: var(--font-weight-regular)
}

._size-2xlarge_1v2ig_13 {
    font-size: var(--font-size-2xl);
    line-height: var(--font-line-height-compact)
}

._size-xlarge_1v2ig_18 {
    font-size: var(--font-size-xl);
    line-height: var(--font-line-height-compact)
}

._size-large_1v2ig_23 {
    font-size: var(--font-size-l);
    line-height: var(--font-line-height-loose)
}

._size-medium_1v2ig_28 {
    font-size: var(--font-size-m);
    line-height: var(--font-line-height-loose)
}

._size-small_1v2ig_33 {
    font-size: var(--font-size-s);
    line-height: var(--font-line-height-regular)
}

._primary_1v2ig_38 {
    color: var(--color-primary)
}

._text-dark_1v2ig_42 {
    color: var(--color-text-dark)
}

._text-base_1v2ig_46 {
    color: var(--color-text-base)
}

._text-light_1v2ig_50 {
    color: var(--color-text-light)
}

._text-xlight_1v2ig_54 {
    color: var(--color-text-xlight)
}

._danger_1v2ig_58 {
    color: var(--color-danger)
}

._align-left_1v2ig_62 {
    text-align: left
}

._align-right_1v2ig_66 {
    text-align: right
}

._align-center_1v2ig_70 {
    text-align: center
}

._callout_z2ivk_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between;
    font-size: var(--font-size-2xs);
    padding: var(--spacing-xs);
    border: var(--border-width-base) var(--border-style-base);
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    line-height: var(--font-line-height-loose)
}

._callout_z2ivk_5._slim_z2ivk_14 {
    line-height: var(--font-line-height-loose);
    padding: var(--spacing-3xs) var(--spacing-2xs)
}

._round_z2ivk_19 {
    border-radius: var(--border-radius-base)
}

._messageSection_z2ivk_23 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._info_z2ivk_28,
._custom_z2ivk_29 {
    border-color: var(--color-foreground-base);
    background-color: var(--color-foreground-xlight);
    color: var(--color-info)
}

._warning_z2ivk_35 {
    border-color: var(--color-warning-tint-1);
    background-color: var(--color-warning-tint-2);
    color: var(--color-warning)
}

._success_z2ivk_41 {
    border-color: var(--color-success-tint-1);
    background-color: var(--color-success-tint-2);
    color: var(--color-success)
}

._danger_z2ivk_47 {
    border-color: var(--color-danger-tint-1);
    background-color: var(--color-danger-tint-2);
    color: var(--color-danger)
}

._icon_z2ivk_53 {
    line-height: 1;
    margin-right: var(--spacing-2xs)
}

._secondary_z2ivk_58 {
    font-size: var(--font-size-2xs);
    font-weight: var(--font-weight-bold);
    color: var(--color-secondary);
    background-color: var(--color-secondary-tint-3);
    border-color: var(--color-secondary-tint-1)
}

._container_g25sr_5 {
    border: 2px dashed var(--color-foreground-base);
    border-radius: var(--border-radius-large);
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding: var(--spacing-3xl)
}

._container_g25sr_5>* {
    margin-bottom: var(--spacing-l)
}

._container_g25sr_5>*:last-child {
    margin-bottom: 0
}

._emoji_g25sr_20 {
    font-size: 40px
}

._heading_g25sr_24 {
    margin-bottom: var(--spacing-l);
    text-align: center
}

._description_g25sr_29 {
    color: var(--color-text-base);
    margin-bottom: var(--spacing-xl);
    text-align: center
}

._callout_g25sr_35 {
    width: 100%;
    text-align: left
}

._activator_4frms_5 {
    cursor: pointer;
    padding: var(--spacing-2xs);
    margin: 0;
    border-radius: var(--border-radius-base);
    line-height: normal !important
}

._activator_4frms_5 svg {
    position: static !important
}

._activator_4frms_5:hover {
    background-color: var(--color-background-base);
    color: initial !important
}

._itemContainer_4frms_20 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex
}

._icon_4frms_24 {
    text-align: center;
    margin-right: var(--spacing-2xs)
}

._icon_4frms_24 svg {
    width: 1.2em !important
}

li.is-disabled ._hasCustomStyling_4frms_32 {
    color: inherit !important
}

._container_13dqp_5>* {
    line-height: 1
}

._button_13dqp_9 {
    cursor: pointer;
    padding: var(--spacing-4xs);
    border-radius: var(--border-radius-base)
}

._button_13dqp_9:hover {
    color: var(--color-primary);
    cursor: pointer
}

._button_13dqp_9:focus {
    color: var(--color-primary)
}

._dark_13dqp_22 {
    color: var(--color-text-dark)
}

._dark_13dqp_22:focus {
    background-color: var(--color-background-xlight)
}

._iconContainer_13dqp_29 {
    display: inline
}

li:hover ._iconContainer_13dqp_29 svg {
    color: var(--color-primary-tint-1)
}

._alert_n2mmu_119 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    position: relative;
    min-height: 60px;
    border-bottom: 1px solid transparent;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between;
    padding: 8px 16px
}

._alert_n2mmu_119._center_n2mmu_128 {
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center
}

._alert_n2mmu_119._success_n2mmu_131._light_n2mmu_131 {
    color: var(--color-success)
}

._alert_n2mmu_119._success_n2mmu_131._light_n2mmu_131._background_n2mmu_134 {
    background-color: hsl(var(--color-success-h), var(--color-success-s), calc(var(--color-success-l) + 48%));
    border-color: var(--color-success)
}

._alert_n2mmu_119._success_n2mmu_131._light_n2mmu_131 ._el-alert__description_n2mmu_138 {
    color: var(--color-success)
}

._alert_n2mmu_119._success_n2mmu_131._dark_n2mmu_141 {
    color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._success_n2mmu_131._dark_n2mmu_141:not(._background_n2mmu_134) {
    color: var(--color-success)
}

._alert_n2mmu_119._success_n2mmu_131._dark_n2mmu_141._background_n2mmu_134 {
    background-color: var(--color-success);
    border-color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._info_n2mmu_151._light_n2mmu_131 {
    color: var(--color-info)
}

._alert_n2mmu_119._info_n2mmu_151._light_n2mmu_131._background_n2mmu_134 {
    background-color: hsl(var(--color-info-h), var(--color-info-s), calc(var(--color-info-l) + 39%));
    border-color: var(--color-info)
}

._alert_n2mmu_119._info_n2mmu_151._dark_n2mmu_141 {
    color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._info_n2mmu_151._dark_n2mmu_141:not(._background_n2mmu_134) {
    color: var(--color-info)
}

._alert_n2mmu_119._info_n2mmu_151._dark_n2mmu_141._background_n2mmu_134 {
    background-color: var(--color-info);
    border-color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._info_n2mmu_151 ._el-alert__description_n2mmu_138 {
    color: var(--color-info)
}

._alert_n2mmu_119._warning_n2mmu_171._light_n2mmu_131 {
    color: var(--color-warning)
}

._alert_n2mmu_119._warning_n2mmu_171._light_n2mmu_131._background_n2mmu_134 {
    background-color: hsl(var(--color-warning-h), var(--color-warning-s), calc(var(--color-warning-l) + 40%));
    border-color: var(--color-warning)
}

._alert_n2mmu_119._warning_n2mmu_171._light_n2mmu_131 ._el-alert__description_n2mmu_138 {
    color: var(--color-warning)
}

._alert_n2mmu_119._warning_n2mmu_171._dark_n2mmu_141 {
    color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._warning_n2mmu_171._dark_n2mmu_141:not(._background_n2mmu_134) {
    color: var(--color-warning)
}

._alert_n2mmu_119._warning_n2mmu_171._dark_n2mmu_141._background_n2mmu_134 {
    background-color: var(--color-warning);
    border-color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._error_n2mmu_191._light_n2mmu_131 {
    color: var(--color-danger)
}

._alert_n2mmu_119._error_n2mmu_191._light_n2mmu_131._background_n2mmu_134 {
    background-color: var(--color-danger-tint-2);
    border-color: var(--color-danger)
}

._alert_n2mmu_119._error_n2mmu_191._light_n2mmu_131 ._el-alert__description_n2mmu_138 {
    color: var(--color-danger)
}

._alert_n2mmu_119._error_n2mmu_191._dark_n2mmu_141 {
    color: var(--color-foreground-xlight)
}

._alert_n2mmu_119._error_n2mmu_191._dark_n2mmu_141:not(._background_n2mmu_134) {
    color: var(--color-danger)
}

._alert_n2mmu_119._error_n2mmu_191._dark_n2mmu_141._background_n2mmu_134 {
    background-color: var(--color-danger);
    border-color: var(--color-foreground-xlight)
}

._content_n2mmu_212 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._icon_n2mmu_217 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    color: inherit;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding-left: var(--spacing-2xs);
    padding-right: var(--spacing-s)
}

._text_n2mmu_225 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center
}

._title_n2mmu_231 {
    font-size: 13px;
    line-height: 18px;
    font-weight: 700
}

._description_n2mmu_237 {
    font-size: 12px
}

._description_n2mmu_237._hasTitle_n2mmu_240 {
    margin: 5px 0 0
}

._aside_n2mmu_244 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding-left: var(--spacing-s)
}

._container_xwcom_5 {
    position: relative;
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._empty_xwcom_12 {
    border-radius: 50%;
    background-color: var(--color-foreground-dark);
    opacity: .3
}

._initials_xwcom_18 {
    position: absolute;
    font-size: var(--font-size-2xs);
    font-weight: var(--font-weight-bold);
    color: var(--color-text-xlight);
    text-shadow: 0px 1px 6px rgba(25, 11, 9, .3)
}

._small_xwcom_26 {
    height: 28px;
    width: 28px
}

._medium_xwcom_31 {
    height: 40px;
    width: 40px
}

._large_xwcom_36 {
    height: 48px;
    width: 48px
}

._badge_11lvf_5 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding: var(--spacing-5xs) var(--spacing-4xs);
    border: var(--border-base);
    white-space: nowrap
}

._default_11lvf_13 {
    border-radius: var(--border-radius-base);
    color: var(--color-text-light);
    border-color: var(--color-text-light)
}

._success_11lvf_20 {
    border-radius: var(--border-radius-base);
    color: var(--color-success);
    border-color: var(--color-success)
}

._warning_11lvf_27 {
    border-radius: var(--border-radius-base);
    color: var(--color-warning);
    border-color: var(--color-warning)
}

._danger_11lvf_34 {
    border-radius: var(--border-radius-base);
    color: var(--color-danger);
    border-color: var(--color-danger)
}

._primary_11lvf_41 {
    padding: var(--spacing-5xs) var(--spacing-3xs);
    border-radius: var(--border-radius-xlarge);
    color: var(--color-foreground-xlight);
    background-color: var(--color-primary);
    border-color: var(--color-primary)
}

._secondary_11lvf_50 {
    border-radius: var(--border-radius-xlarge);
    color: var(--color-secondary);
    background-color: var(--color-secondary-tint-1)
}

._tertiary_11lvf_57 {
    border-radius: var(--border-radius-base);
    color: var(--color-text-light);
    border-color: var(--color-text-light);
    padding: 1px var(--spacing-5xs)
}

._uiBlocker_1x8fa_5 {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: var(--color-background-dark);
    z-index: 10;
    opacity: .6;
    border-radius: var(--border-radius-large)
}

[data-v-87c80631]:root {
    --header-height: 65
}

.fade-enter-active[data-v-87c80631],
.fade-leave-active[data-v-87c80631] {
    -webkit-transition: opacity .2s;
    transition: opacity .2s
}

.fade-enter[data-v-87c80631],
.fade-leave-to[data-v-87c80631] {
    opacity: 0
}

._card_e3hfs_5 {
    border-radius: var(--border-radius-large);
    border: var(--border-base);
    background-color: var(--color-background-xlight);
    padding: var(--spacing-s);
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: horizontal;
    -webkit-box-direction: normal;
    -ms-flex-direction: row;
    flex-direction: row;
    width: 100%;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._header_e3hfs_16,
._footer_e3hfs_17 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: horizontal;
    -webkit-box-direction: normal;
    -ms-flex-direction: row;
    flex-direction: row;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._content_e3hfs_24 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between;
    -webkit-box-flex: 1;
    -ms-flex: 1;
    flex: 1;
    width: 100%
}

._body_e3hfs_32 {
    width: 100%;
    height: 100%;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between
}

._icon_e3hfs_40 {
    width: 24px;
    height: 24px;
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    margin-right: var(--spacing-s)
}

._hoverable_e3hfs_49 {
    cursor: pointer;
    -webkit-transition-property: border, color;
    transition-property: border, color;
    -webkit-transition-duration: .3s;
    transition-duration: .3s;
    -webkit-transition-timing-function: ease;
    transition-timing-function: ease
}

._hoverable_e3hfs_49:hover,
._hoverable_e3hfs_49:focus {
    color: var(--color-primary);
    border-color: var(--color-primary)
}

._xlarge_14brn_5 {
    --input-font-size: var(--font-size-m)
}

._xlarge_14brn_5 input {
    height: 48px
}

._limitPopperWidth_14brn_12 {
    width: 0
}

._limitPopperWidth_14brn_12 li>span {
    text-overflow: ellipsis;
    overflow-x: hidden
}

._container_14brn_20 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    width: 100%
}

._withPrepend_14brn_25 input {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0
}

@-moz-document url-prefix() {
    ._withPrepend_14brn_25 input {
        padding: 0 var(--spacing-3xs)
    }
}

._prepend_14brn_35 {
    font-size: var(--font-size-2xs);
    border: var(--border-base);
    border-right: none;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    padding: 0 var(--spacing-3xs);
    background-color: var(--color-background-light);
    border-bottom-left-radius: var(--input-border-radius, var(--border-radius-base));
    border-top-left-radius: var(--input-border-radius, var(--border-radius-base));
    color: var(--color-text-base);
    white-space: nowrap
}

._datatableWrapper_1fhod_5 {
    display: block;
    width: 100%
}

._datatable_1fhod_5 {
    width: 100%
}

._datatable_1fhod_5 tbody tr td {
    vertical-align: top;
    color: var(--color-text-base);
    padding: var(--spacing-s) var(--spacing-2xs)
}

._datatable_1fhod_5 tbody tr:nth-of-type(even) {
    background: var(--color-background-xlight)
}

._datatable_1fhod_5 tbody tr:nth-of-type(odd) {
    background: var(--color-background-light)
}

._datatableHeader_1fhod_25 {
    background: var(--color-background-base)
}

._datatableHeader_1fhod_25 th {
    text-align: left;
    padding: var(--spacing-s) var(--spacing-2xs)
}

._pagination_1fhod_33 {
    width: 100%;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    bottom: 0;
    overflow: visible;
    margin-top: var(--spacing-s)
}

._pageSizeSelector_1fhod_43 {
    text-transform: capitalize;
    max-width: 150px;
    -webkit-box-flex: 0;
    -ms-flex: 0 1 auto;
    flex: 0 1 auto
}

._xlarge_xj2wx_5 {
    --input-font-size: var(--font-size-m)
}

._xlarge_xj2wx_5 input {
    height: 48px
}

._buttons_1fogc_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    margin-top: var(--spacing-s)
}

._container_1woiu_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column
}

._inputLabel_1woiu_10 {
    display: block
}

._container_1woiu_5:hover ._infoIcon_1woiu_14,
._inputLabel_1woiu_10:hover ._infoIcon_1woiu_14 {
    opacity: 1
}

._container_1woiu_5:hover ._options_1woiu_18,
._inputLabel_1woiu_10:hover ._options_1woiu_18,
._container_1woiu_5:hover ._overlay_1woiu_23,
._inputLabel_1woiu_10:hover ._overlay_1woiu_23 {
    opacity: 1;
    -webkit-transition: opacity .1s ease-in;
    transition: opacity .1s ease-in
}

._title_1woiu_29 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    min-width: 0
}

._title_1woiu_29>* {
    white-space: nowrap
}

._infoIcon_1woiu_14 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    color: var(--color-text-light);
    padding-left: var(--spacing-4xs);
    background-color: var(--color-background-xlight);
    z-index: 1
}

._options_1woiu_18 {
    opacity: 0;
    background-color: var(--color-background-xlight);
    -webkit-transition: opacity .25s cubic-bezier(.98, -.06, .49, -.2);
    transition: opacity .25s cubic-bezier(.98, -.06, .49, -.2)
}

._options_1woiu_18>* {
    float: right
}

._overlay_1woiu_23 {
    position: relative;
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1;
    opacity: 0;
    -webkit-transition: opacity .25s cubic-bezier(.98, -.06, .49, -.2);
    transition: opacity .25s cubic-bezier(.98, -.06, .49, -.2)
}

._overlay_1woiu_23>div {
    position: absolute;
    width: 60px;
    height: 19px;
    top: 0;
    right: 0;
    z-index: 0;
    background: -webkit-gradient(linear, right top, left top, color-stop(72.19%, var(--color-foreground-xlight)), color-stop(107.45%, rgba(255, 255, 255, 0)));
    background: linear-gradient(270deg, var(--color-foreground-xlight) 72.19%, rgba(255, 255, 255, 0) 107.45%)
}

._hidden_1woiu_72 {
    opacity: 0
}

._visible_1woiu_76 {
    opacity: 1
}

._heading_1woiu_80 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex
}

._overflow_1woiu_84 {
    overflow-x: hidden;
    overflow-y: clip
}

._small_1woiu_89 {
    margin-bottom: var(--spacing-5xs)
}

._medium_1woiu_93 {
    margin-bottom: var(--spacing-2xs)
}

._underline_1woiu_97 {
    border-bottom: var(--border-base)
}

:root ._tooltipPopper_1woiu_101 {
    max-width: 400px
}

:root ._tooltipPopper_1woiu_101 li {
    margin-left: var(--spacing-s)
}

._n8nCheckbox_1i6bx_5 {
    display: -webkit-box !important;
    display: -ms-flexbox !important;
    display: flex !important;
    white-space: normal !important;
    margin-bottom: var(--spacing-2xs)
}

._n8nCheckbox_1i6bx_5 span {
    white-space: normal
}

._n8nCheckbox_1i6bx_5 label {
    cursor: pointer;
    margin-bottom: 0
}

._infoText_z4p5l_5 {
    margin-top: var(--spacing-2xs);
    font-size: var(--font-size-2xs);
    font-weight: var(--font-weight-regular);
    color: var(--color-text-base)
}

._errors_z4p5l_12 {
    color: var(--color-danger)
}

._errorInput_z4p5l_17 {
    --input-border-color: var(--color-danger)
}

._grid_c0i81_5 {
    display: grid;
    grid-row-gap: var(--spacing-s);
    grid-column-gap: var(--spacing-2xs)
}

._gridMulti_c0i81_11 {
    grid-template-columns: repeat(2, 1fr)
}

._primary_omopq_5 {
    color: var(--color-primary)
}

._primary_omopq_5:active {
    color: hsl(var(--color-primary-h), calc(var(--color-primary-s) + -30%), var(--color-primary-l))
}

._text_omopq_12 {
    color: var(--color-text-base)
}

._text_omopq_12:hover {
    color: var(--color-primary)
}

._text_omopq_12:active {
    color: hsl(var(--color-primary-h), calc(var(--color-primary-s) + -30%), var(--color-primary-l))
}

._danger_omopq_22 {
    color: var(--color-danger)
}

._danger_omopq_22:active {
    color: hsl(var(--color-danger-h), calc(var(--color-danger-s) + -20%), var(--color-danger-l))
}

._secondary_omopq_29:active {
    color: hsl(var(--color-secondary-h), calc(var(--color-secondary-s) + -20%), var(--color-secondary-l))
}

._secondary_omopq_29 {
    color: var(--color-secondary)
}

._primary-underline_omopq_40,
._text-underline_omopq_45,
._danger-underline_omopq_50,
._secondary-underline_omopq_55 {
    text-decoration: underline
}

._heading_1dsc6_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    margin-bottom: var(--spacing-xl)
}

._container_1dsc6_11 {
    background-color: var(--color-background-xlight);
    padding: var(--spacing-l);
    border: var(--border-base);
    border-radius: var(--border-radius-large);
    -webkit-box-shadow: 0px 4px 16px rgba(99, 77, 255, .06);
    box-shadow: 0 4px 16px #634dff0f
}

._inputsContainer_1dsc6_19 {
    margin-bottom: var(--spacing-xl)
}

._actionContainer_1dsc6_23 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center
}

._buttonsContainer_1dsc6_28 {
    margin-bottom: var(--spacing-s)
}

._secondaryButtonContainer_1dsc6_33 {
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._withLabel_1dsc6_39 {
    margin-bottom: var(--spacing-s)
}

._container_zzhlv_5 {
    background-color: var(--color-background-base)
}

._header_zzhlv_9 {
    cursor: pointer;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    padding: var(--spacing-s);
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._header_zzhlv_9 ._headerText_zzhlv_15 {
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1
}

._expanded_zzhlv_19 {
    padding: var(--spacing-s) var(--spacing-s) var(--spacing-2xs) var(--spacing-s)
}

._accordionItems_zzhlv_23 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical !important;
    -webkit-box-direction: normal !important;
    -ms-flex-direction: column !important;
    flex-direction: column !important;
    -webkit-box-align: start !important;
    -ms-flex-align: start !important;
    align-items: flex-start !important;
    width: 100%
}

._accordionItem_zzhlv_23 {
    display: block !important;
    text-align: left
}

._description_zzhlv_35 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    padding: 0 var(--spacing-s) var(--spacing-s) var(--spacing-s)
}

._description_zzhlv_35 b {
    font-weight: var(--font-weight-bold)
}

._base_w9w37_5 {
    font-size: var(--font-size-2xs);
    line-height: var(--font-size-s);
    word-break: normal;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._base_w9w37_5 svg {
    font-size: var(--font-size-s)
}

._bold_w9w37_16 {
    font-weight: var(--font-weight-bold)
}

._note_w9w37_20 svg {
    margin-right: var(--spacing-4xs)
}

._tooltip_w9w37_27 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex
}

._iconText_w9w37_32 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: start;
    -ms-flex-align: start;
    align-items: flex-start
}

._info-light_w9w37_37 {
    color: var(--color-foreground-dark)
}

._info_w9w37_37 {
    color: var(--color-text-light)
}

._warning_w9w37_45 {
    color: var(--color-warning)
}

._danger_w9w37_49 {
    color: var(--color-danger)
}

._h1Last_bim62_5 {
    width: 40%
}

._pLast_bim62_9 {
    width: 61%
}

._custom_bim62_13 {
    width: 100%;
    height: 100%
}

.n8n-loading-custom .el-skeleton,
.n8n-loading-custom .el-skeleton .el-skeleton__item {
    width: 100%;
    height: 100%
}

._markdown_1i6uh_5 {
    color: var(--color-text-base)
}

._markdown_1i6uh_5 * {
    font-size: var(--font-size-m);
    line-height: var(--font-line-height-xloose)
}

._markdown_1i6uh_5 h1,
._markdown_1i6uh_5 h2,
._markdown_1i6uh_5 h3,
._markdown_1i6uh_5 h4 {
    margin-bottom: var(--spacing-s);
    font-size: var(--font-size-m);
    font-weight: var(--font-weight-bold)
}

._markdown_1i6uh_5 h3,
._markdown_1i6uh_5 h4 {
    font-weight: var(--font-weight-bold)
}

._markdown_1i6uh_5 p,
._markdown_1i6uh_5 span {
    margin-bottom: var(--spacing-s)
}

._markdown_1i6uh_5 ul,
._markdown_1i6uh_5 ol {
    margin-bottom: var(--spacing-s);
    padding-left: var(--spacing-m)
}

._markdown_1i6uh_5 ul li,
._markdown_1i6uh_5 ol li {
    margin-top: .25em
}

._markdown_1i6uh_5 pre {
    margin-bottom: var(--spacing-s);
    display: grid
}

._markdown_1i6uh_5 pre>code {
    display: block;
    padding: var(--spacing-s);
    color: var(--color-text-dark);
    background-color: var(--color-background-base);
    overflow-x: auto
}

._markdown_1i6uh_5 li>code,
._markdown_1i6uh_5 p>code {
    padding: 0 var(--spacing-4xs);
    color: var(--color-text-dark);
    background-color: var(--color-background-base)
}

._markdown_1i6uh_5 ._label_1i6uh_54 {
    color: var(--color-text-base)
}

._markdown_1i6uh_5 img {
    width: 100%;
    max-height: 90vh;
    -o-object-fit: cover;
    object-fit: cover;
    border: var(--border-width-base) var(--color-foreground-base) var(--border-style-base);
    border-radius: var(--border-radius-large)
}

._markdown_1i6uh_5 blockquote {
    padding-left: 10px;
    font-style: italic;
    border-left: var(--border-color-base) 2px solid
}

._sticky_1i6uh_70 {
    color: var(--color-text-dark)
}

._sticky_1i6uh_70 h1,
._sticky_1i6uh_70 h2,
._sticky_1i6uh_70 h3,
._sticky_1i6uh_70 h4 {
    margin-bottom: var(--spacing-2xs);
    font-weight: var(--font-weight-bold);
    line-height: var(--font-line-height-loose)
}

._sticky_1i6uh_70 h1 {
    font-size: 36px
}

._sticky_1i6uh_70 h2 {
    font-size: 24px
}

._sticky_1i6uh_70 h3,
._sticky_1i6uh_70 h4,
._sticky_1i6uh_70 h5,
._sticky_1i6uh_70 h6 {
    font-size: var(--font-size-m)
}

._sticky_1i6uh_70 p {
    margin-bottom: var(--spacing-2xs);
    font-size: var(--font-size-s);
    font-weight: var(--font-weight-regular);
    line-height: var(--font-line-height-loose)
}

._sticky_1i6uh_70 ul,
._sticky_1i6uh_70 ol {
    margin-bottom: var(--spacing-2xs);
    padding-left: var(--spacing-m)
}

._sticky_1i6uh_70 ul li,
._sticky_1i6uh_70 ol li {
    margin-top: .25em;
    font-size: var(--font-size-s);
    font-weight: var(--font-weight-regular);
    line-height: var(--font-line-height-regular)
}

._sticky_1i6uh_70 code {
    background-color: var(--color-background-base);
    padding: 0 var(--spacing-4xs);
    color: var(--color-secondary)
}

._sticky_1i6uh_70 pre>code,
._sticky_1i6uh_70 li>code,
._sticky_1i6uh_70 p>code {
    color: var(--color-secondary)
}

._sticky_1i6uh_70 a:hover {
    text-decoration: underline
}

._sticky_1i6uh_70 img {
    -o-object-fit: contain;
    object-fit: contain
}

._sticky_1i6uh_70 img[src*="#full-width"] {
    width: 100%
}

._spacer_1i6uh_131 {
    margin: var(--spacing-2xl)
}

.el-menu-item,
.el-submenu__title {
    --menu-font-color: var(--color-text-base);
    --menu-item-active-background-color: var(--color-foreground-base);
    --menu-item-active-font-color: var(--color-text-dark);
    --menu-item-hover-fill: var(--color-foreground-base);
    --menu-item-hover-font-color: var(--color-text-dark);
    --menu-item-height: 35px;
    --submenu-item-height: 27px
}

._submenu_1h6xl_16 {
    background: none !important
}

._submenu_1h6xl_16._compact_1h6xl_19 .el-submenu__title i {
    display: none
}

._submenu_1h6xl_16 .el-submenu__title {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    border-radius: var(--border-radius-base) !important;
    padding: var(--spacing-2xs) var(--spacing-xs) !important;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}

._submenu_1h6xl_16 .el-submenu__title i {
    padding-top: 2px
}

._submenu_1h6xl_16 .el-submenu__title i:hover {
    color: var(--color-primary)
}

._submenu_1h6xl_16 .el-submenu__title:hover ._icon_1h6xl_35 {
    color: var(--color-text-dark)
}

._submenu_1h6xl_16 ._menuItem_1h6xl_38 {
    height: var(--submenu-item-height) !important;
    min-width: auto !important;
    margin: var(--spacing-2xs) 0 !important;
    padding-left: var(--spacing-l) !important;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}

._submenu_1h6xl_16 ._menuItem_1h6xl_38:hover ._icon_1h6xl_35 {
    color: var(--color-text-dark)
}

._disableActiveStyle_1h6xl_49 {
    background-color: initial !important;
    color: var(--color-text-base) !important
}

._disableActiveStyle_1h6xl_49 svg {
    color: var(--color-text-base) !important
}

._disableActiveStyle_1h6xl_49:hover {
    background-color: var(--color-foreground-base) !important
}

._disableActiveStyle_1h6xl_49:hover svg {
    color: var(--color-text-dark) !important
}

._disableActiveStyle_1h6xl_49:hover.el-submenu {
    background-color: unset !important
}

._active_1h6xl_66,
._active_1h6xl_66 .el-submenu__title {
    background-color: var(--color-foreground-base);
    border-radius: var(--border-radius-base)
}

._active_1h6xl_66 ._icon_1h6xl_35,
._active_1h6xl_66 .el-submenu__title ._icon_1h6xl_35 {
    color: var(--color-text-dark)
}

._menuItem_1h6xl_38 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    padding: var(--spacing-2xs) var(--spacing-xs) !important;
    margin: 0 !important;
    border-radius: var(--border-radius-base) !important;
    overflow: hidden
}

._icon_1h6xl_35 {
    min-width: var(--spacing-s);
    margin-right: var(--spacing-xs);
    text-align: center
}

._secondaryIcon_1h6xl_88 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    -webkit-box-pack: end;
    -ms-flex-pack: end;
    justify-content: flex-end;
    -webkit-box-flex: 1;
    -ms-flex: 1;
    flex: 1
}

._label_1h6xl_95 {
    overflow: hidden;
    text-overflow: ellipsis;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}

._item_1h6xl_101+._item_1h6xl_101 {
    margin-top: 8px !important
}

._compact_1h6xl_19 {
    width: 40px
}

._compact_1h6xl_19 ._icon_1h6xl_35 {
    margin: 0;
    overflow: visible !important;
    visibility: visible !important;
    width: initial !important;
    height: initial !important
}

._compact_1h6xl_19 ._label_1h6xl_95,
._compact_1h6xl_19 ._secondaryIcon_1h6xl_88 {
    display: none
}

._submenuPopper_1h6xl_122 {
    display: block
}

._submenuPopper_1h6xl_122 ul {
    padding: 0 var(--spacing-xs) !important
}

._submenuPopper_1h6xl_122 ._menuItem_1h6xl_38 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    padding: var(--spacing-2xs) var(--spacing-xs) !important;
    margin: var(--spacing-2xs) 0 !important
}

._submenuPopper_1h6xl_122 ._icon_1h6xl_35 {
    margin-right: var(--spacing-xs)
}

._submenuPopper_1h6xl_122 ._label_1h6xl_95 {
    display: block
}

._container_1gao6_5 {
    height: 100%;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    background-color: var(--menu-background, var(--color-background-xlight))
}

._menuContent_1gao6_12 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-pack: justify;
    -ms-flex-pack: justify;
    justify-content: space-between;
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1
}

._menuContent_1gao6_12>div>.el-menu {
    background: none;
    padding: var(--menu-padding, 12px)
}

._upperContent_1gao6_23 ul {
    padding-top: 0 !important
}

._lowerContent_1gao6_27 ul {
    padding-bottom: 0 !important
}

._menuCollapsed_1gao6_31 {
    -webkit-transition: width .15s ease-in-out;
    transition: width .15s ease-in-out
}

._menuCollapsed_1gao6_31 .hideme {
    display: none !important
}

._creatorNode_qpkac_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    cursor: pointer;
    z-index: 1;
    padding: var(--spacing-xs) var(--spacing-2xs) var(--spacing-xs) 0
}

._creatorNode_qpkac_5._hasAction_qpkac_12 {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}

._creatorNode_qpkac_5:hover ._panelIcon_qpkac_16 {
    color: var(--action-arrow-color-hover, var(--color-text-light))
}

._panelIcon_qpkac_16 {
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: end;
    -ms-flex-pack: end;
    justify-content: flex-end;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    margin-left: var(--spacing-2xs);
    color: var(--action-arrow-color, var(--color-text-lighter));
    cursor: pointer;
    background: transparent;
    border: none
}

._tooltipIcon_qpkac_32 {
    margin-left: var(--spacing-3xs)
}

._panelArrow_qpkac_36 {
    font-size: var(--font-size-2xs);
    width: 12px
}

._details_qpkac_41 {
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._nodeIcon_qpkac_45 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    margin-right: var(--node-icon-margin-right, var(--spacing-s))
}

._name_qpkac_50 {
    font-weight: var(--node-creator-name-weight, var(--font-weight-bold));
    font-size: var(--node-creator-name-size, var(--font-size-s));
    line-height: 1.115rem
}

._description_qpkac_56 {
    margin-top: var(--spacing-5xs);
    font-size: var(--font-size-2xs);
    line-height: 1rem;
    font-weight: 400;
    color: var(--node-creator-description-colos, var(--color-text-base))
}

._triggerIcon_qpkac_64 {
    margin-left: var(--spacing-3xs);
    color: var(--color-primary)
}

[data-v-64dd9238]:root {
    --header-height: 65
}

.el-tooltip svg[data-v-64dd9238] {
    color: var(--color-foreground-xdark)
}

._nodeIconWrapper_1nau5_5 {
    width: var(--node-icon-size, 26px);
    height: var(--node-icon-size, 26px);
    border-radius: var(--border-radius-small);
    color: var(--node-icon-color, #444);
    line-height: var(--node-icon-size, 26px);
    font-size: 1.1em;
    overflow: hidden;
    text-align: center;
    font-weight: 700;
    font-size: 20px
}

._icon_1nau5_18 {
    height: 100%;
    width: 100%;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    pointer-events: none
}

._icon_1nau5_18 svg {
    max-width: 100%;
    max-height: 100%
}

._nodeIconPlaceholder_1nau5_31 {
    text-align: center
}

._nodeIconImage_1nau5_35 {
    width: 100%;
    max-width: 100%;
    max-height: 100%
}

._circle_1nau5_41 {
    border-radius: 50%
}

._disabled_1nau5_45 {
    color: "#ccc";
    -webkit-filter: contrast(40%) brightness(1.5) grayscale(100%);
    filter: contrast(40%) brightness(1.5) grayscale(100%)
}

._notice_1659n_5 {
    font-size: var(--font-size-2xs);
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    color: var(--custom-font-black);
    margin: var(--notice-margin, var(--spacing-s) 0);
    padding: var(--spacing-2xs);
    background-color: var(--background-color);
    border-width: 1px 1px 1px 7px;
    border-style: solid;
    border-color: var(--border-color);
    border-radius: var(--border-radius-small);
    line-height: var(--font-line-height-compact)
}

._notice_1659n_5 a {
    font-weight: var(--font-weight-bold)
}

._warning_1659n_22 {
    --border-color: var(--color-warning-tint-1);
    --background-color: var(--color-warning-tint-2)
}

._danger_1659n_27 {
    --border-color: var(--color-danger-tint-1);
    --background-color: var(--color-danger-tint-2)
}

._success_1659n_32 {
    --border-color: var(--color-success-tint-1);
    --background-color: var(--color-success-tint-2)
}

._info_1659n_37 {
    --border-color: var(--color-info-tint-1);
    --background-color: var(--color-info-tint-2)
}

._expanded_1659n_42+span {
    margin-top: var(--spacing-4xs);
    display: block
}

._truncated_1659n_47 {
    display: inline
}

._pulseContainer_1g8nb_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    height: 240px;
    width: 100%
}

._pulse_1g8nb_5 {
    width: 74px;
    height: 74px;
    border-radius: 50%;
    -webkit-box-shadow: 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
    box-shadow: 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
    -webkit-animation: _pulse_1g8nb_5 6s infinite cubic-bezier(.33, 1, .68, 1);
    animation: _pulse_1g8nb_5 6s infinite cubic-bezier(.33, 1, .68, 1)
}

._pulse2_1g8nb_21 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    width: 74px;
    height: 74px;
    border-radius: 50%;
    -webkit-box-shadow: 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
    box-shadow: 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
    -webkit-animation: _pulse2_1g8nb_21 6s infinite cubic-bezier(.33, 1, .68, 1);
    animation: _pulse2_1g8nb_21 6s infinite cubic-bezier(.33, 1, .68, 1)
}

@-webkit-keyframes _pulse_1g8nb_5 {
    0% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    58.33% {
        -webkit-box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0);
        box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0)
    }

    66.6% {
        -webkit-box-shadow: 0 0 0 66px transparent;
        box-shadow: 0 0 0 66px transparent
    }

    66.7% {
        -webkit-box-shadow: 0 0 0 0 transparent;
        box-shadow: 0 0 0 0 transparent
    }
}

@keyframes _pulse_1g8nb_5 {
    0% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    58.33% {
        -webkit-box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0);
        box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0)
    }

    66.6% {
        -webkit-box-shadow: 0 0 0 66px transparent;
        box-shadow: 0 0 0 66px transparent
    }

    66.7% {
        -webkit-box-shadow: 0 0 0 0 transparent;
        box-shadow: 0 0 0 0 transparent
    }
}

@-webkit-keyframes _pulse2_1g8nb_21 {
    0% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    16.66% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    50% {
        -webkit-box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0);
        box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0)
    }

    83.3% {
        -webkit-box-shadow: 0 0 0 66px transparent;
        box-shadow: 0 0 0 66px transparent
    }

    83.4% {
        -webkit-box-shadow: 0 0 0 0 transparent;
        box-shadow: 0 0 0 0 transparent
    }
}

@keyframes _pulse2_1g8nb_21 {
    0% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    16.66% {
        -webkit-box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4);
        box-shadow: 0 0 0 0 hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), .4)
    }

    50% {
        -webkit-box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0);
        box-shadow: 0 0 0 60px hsla(var(--color-primary-h), var(--color-primary-s), var(--color-primary-l), 0)
    }

    83.3% {
        -webkit-box-shadow: 0 0 0 66px transparent;
        box-shadow: 0 0 0 66px transparent
    }

    83.4% {
        -webkit-box-shadow: 0 0 0 0 transparent;
        box-shadow: 0 0 0 0 transparent
    }
}

._container_zf0sd_5 {
    display: inline-block;
    outline: 0;
    position: relative
}

._hoverable_zf0sd_11:hover ._button_zf0sd_11:not(._active_zf0sd_11) {
    color: var(--color-primary)
}

._input_zf0sd_15 {
    opacity: 0;
    outline: 0;
    z-index: -1;
    position: absolute
}

._button_zf0sd_11 {
    border-radius: 0;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    border-radius: var(--border-radius-base);
    font-weight: var(--font-weight-bold);
    color: var(--color-text-base);
    -webkit-transition: background-color .2s ease;
    transition: background-color .2s ease;
    cursor: pointer;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}

._disabled_zf0sd_34 {
    cursor: not-allowed
}

._medium_zf0sd_38 {
    height: 26px;
    font-size: var(--font-size-2xs);
    padding: 0 var(--spacing-xs)
}

._small_zf0sd_44 {
    font-size: var(--font-size-3xs);
    height: 15px;
    padding: 0 var(--spacing-4xs)
}

._active_zf0sd_11 {
    background-color: var(--color-foreground-xlight);
    color: var(--color-text-dark)
}

._radioGroup_1kp8c_5 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    line-height: 1;
    vertical-align: middle;
    font-size: 0;
    background-color: var(--color-foreground-base);
    padding: var(--spacing-5xs);
    border-radius: var(--border-radius-base)
}

._disabled_1kp8c_15 {
    cursor: not-allowed
}

._resize_1d9yl_5 {
    position: relative;
    width: 100%;
    height: 100%;
    z-index: 2
}

._resizer_1d9yl_12 {
    position: absolute;
    z-index: 3
}

._right_1d9yl_17 {
    width: 12px;
    height: 100%;
    top: -2px;
    right: -2px;
    cursor: ew-resize
}

._top_1d9yl_25 {
    width: 100%;
    height: 12px;
    top: -2px;
    left: -2px;
    cursor: ns-resize
}

._bottom_1d9yl_33 {
    width: 100%;
    height: 12px;
    bottom: -2px;
    left: -2px;
    cursor: ns-resize
}

._left_1d9yl_41 {
    width: 12px;
    height: 100%;
    top: -2px;
    left: -2px;
    cursor: ew-resize
}

._topLeft_1d9yl_49 {
    width: 12px;
    height: 12px;
    top: -3px;
    left: -3px;
    cursor: nw-resize
}

._topRight_1d9yl_57 {
    width: 12px;
    height: 12px;
    top: -3px;
    right: -3px;
    cursor: ne-resize
}

._bottomLeft_1d9yl_65 {
    width: 12px;
    height: 12px;
    bottom: -3px;
    left: -3px;
    cursor: sw-resize
}

._bottomRight_1d9yl_73 {
    width: 12px;
    height: 12px;
    bottom: -3px;
    right: -3px;
    cursor: se-resize
}

._sticky_aan73_5 {
    position: absolute;
    background-color: var(--color-sticky-default-background);
    border: 1px solid var(--color-sticky-default-border);
    border-radius: var(--border-radius-base)
}

._clickable_aan73_12 {
    cursor: pointer
}

._wrapper_aan73_16 {
    width: 100%;
    height: 100%;
    position: absolute;
    padding: var(--spacing-2xs) var(--spacing-xs) 0;
    overflow: hidden
}

._wrapper_aan73_16:after {
    content: "";
    width: 100%;
    height: 24px;
    left: 0;
    bottom: 0;
    position: absolute;
    background: -webkit-gradient(linear, left top, left bottom, from(var(--color-sticky-default-background)), color-stop(.01%, rgba(255, 245, 214, 0)), to(var(--color-sticky-default-background)));
    background: linear-gradient(180deg, var(--color-sticky-default-background), rgba(255, 245, 214, 0) .01%, var(--color-sticky-default-background));
    border-radius: var(--border-radius-base)
}

._footer_aan73_34 {
    padding: var(--spacing-5xs) var(--spacing-2xs) 0 var(--spacing-2xs);
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-pack: end;
    -ms-flex-pack: end;
    justify-content: flex-end
}

.sticky-textarea {
    height: calc(100% - var(--spacing-l));
    padding: var(--spacing-2xs) var(--spacing-2xs) 0 var(--spacing-2xs);
    cursor: default
}

.sticky-textarea .el-textarea {
    height: 100%
}

.sticky-textarea .el-textarea .el-textarea__inner {
    height: 100%;
    resize: unset
}

.full-height {
    height: calc(100% - var(--spacing-2xs))
}

._container_1pbqs_5 {
    position: relative;
    height: 24px;
    min-height: 24px;
    width: 100%
}

._tabs_1pbqs_12 {
    color: var(--color-text-base);
    font-weight: var(--font-weight-bold);
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    width: 100%;
    position: absolute;
    overflow-x: scroll;
    -ms-overflow-style: none;
    scrollbar-width: none
}

._tabs_1pbqs_12::-webkit-scrollbar {
    display: none
}

._tab_1pbqs_12 {
    display: block;
    padding: 0 var(--spacing-s) var(--spacing-2xs) var(--spacing-s);
    padding-bottom: var(--spacing-2xs);
    font-size: var(--font-size-s);
    cursor: pointer;
    white-space: nowrap
}

._tab_1pbqs_12:hover {
    color: var(--color-primary)
}

._activeTab_1pbqs_40 {
    color: var(--color-primary);
    border-bottom: var(--color-primary) 2px solid
}

._alignRight_1pbqs_45 {
    margin-left: auto
}

._link_1pbqs_49 {
    cursor: pointer;
    color: var(--color-text-base)
}

._link_1pbqs_49:hover {
    color: var(--color-primary)
}

._link_1pbqs_49:hover ._external_1pbqs_56 {
    display: inline-block
}

._external_1pbqs_56 {
    display: none
}

._button_1pbqs_64 {
    position: absolute;
    background-color: var(--color-background-base);
    z-index: 1;
    height: 24px;
    width: 10px;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    font-weight: var(--font-weight-bold)
}

._back_1pbqs_75 {
    left: 0
}

._next_1pbqs_81 {
    right: 0
}

._tag_d0kvy_5 {
    min-width: -webkit-max-content;
    min-width: -moz-max-content;
    min-width: max-content;
    padding: 1px var(--spacing-4xs);
    color: var(--color-text-dark);
    background-color: var(--color-background-base);
    border-radius: var(--border-radius-base);
    font-size: var(--font-size-2xs);
    cursor: pointer;
    -webkit-transition: background-color .3s ease;
    transition: background-color .3s ease
}

._tag_d0kvy_5:hover {
    background-color: hsl(var(--color-background-base-h), var(--color-background-base-s), calc(var(--color-background-base-l) - 4%))
}

._tags_h85of_5 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -ms-flex-wrap: wrap;
    flex-wrap: wrap;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    overflow-x: scroll;
    -ms-overflow-style: none;
    scrollbar-width: none;
    margin-top: calc(var(--spacing-4xs) * -1)
}

._tags_h85of_5::-webkit-scrollbar {
    display: none
}

._tags_h85of_5 * {
    margin: var(--spacing-4xs) var(--spacing-4xs) 0 0
}

._indent_ze7xv_5 {
    margin-left: var(--spacing-s)
}

._simple_ze7xv_9 {
    text-indent: calc(var(--spacing-s) * -1);
    margin-left: var(--spacing-s);
    max-width: 300px
}

._container_16jaa_5 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    overflow: hidden
}

._avatarContainer_16jaa_10 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    color: var(--color-text-light)
}

._infoContainer_16jaa_17 {
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1;
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-orient: vertical;
    -webkit-box-direction: normal;
    -ms-flex-direction: column;
    flex-direction: column;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    margin-left: var(--spacing-xs)
}

._pendingUser_16jaa_25 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    justify-content: center;
    margin-left: var(--spacing-xs)
}

._pendingBadge_16jaa_32 {
    margin-left: var(--spacing-xs)
}

._disabled_16jaa_36 {
    opacity: .5
}

._itemContainer_18yo2_5 {
    --select-option-padding: var(--spacing-2xs) var(--spacing-s);
    --select-option-line-height: 1
}

:root ._limitPopperWidth_18yo2_10 {
    width: 0
}

:root ._limitPopperWidth_18yo2_10 li>span {
    text-overflow: ellipsis;
    overflow-x: hidden
}

._itemContainer_izbat_5 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    padding: var(--spacing-2xs) 0 vaR(--spacing-2xs) 0
}

._itemContainer_izbat_5>*:first-child {
    -webkit-box-flex: 1;
    -ms-flex-positive: 1;
    flex-grow: 1
}

._itemWithBorder_izbat_13 {
    border-bottom: var(--border-base)
}

._badgeContainer_izbat_18 {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._badgeContainer_izbat_18>* {
    margin-left: var(--spacing-2xs)
}

.recycle-scroller-wrapper {
    height: 100%;
    width: 100%;
    overflow: auto;
    -webkit-box-flex: 1;
    -ms-flex: 1 1 auto;
    flex: 1 1 auto
}

.recycle-scroller {
    width: 100%;
    display: block;
    position: relative
}

.recycle-scroller-items-wrapper {
    position: absolute;
    width: 100%
}

.recycle-scroller-item {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    position: relative;
    width: 100%
}

:root {
    --header-height: 65
}

._component_1am9e_5 {
    display: -webkit-inline-box;
    display: -ms-inline-flexbox;
    display: inline-flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center
}

._input_1am9e_10 {
    margin-left: var(--spacing-3xs)
}

[data-v-e550a5be]:root {
    --header-height: 65
}

[data-v-e550a5be] .el-color-picker .el-color-picker__empty,
[data-v-e550a5be] .el-color-picker .el-color-picker__icon {
    display: none
}
  /* You can include the styles here, or import from an external CSS file */
  </style>
  