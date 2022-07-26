// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Review
{
    [HerokuResourceType("heroku:review/appConfig:AppConfig")]
    public partial class AppConfig : global::Pulumi.CustomResource
    {
        [Output("automaticReviewApps")]
        public Output<bool?> AutomaticReviewApps { get; private set; } = null!;

        [Output("baseName")]
        public Output<string> BaseName { get; private set; } = null!;

        [Output("deployTarget")]
        public Output<Outputs.AppConfigDeployTarget> DeployTarget { get; private set; } = null!;

        [Output("destroyStaleApps")]
        public Output<bool?> DestroyStaleApps { get; private set; } = null!;

        [Output("orgRepo")]
        public Output<string> OrgRepo { get; private set; } = null!;

        [Output("pipelineId")]
        public Output<string> PipelineId { get; private set; } = null!;

        [Output("repoId")]
        public Output<int> RepoId { get; private set; } = null!;

        [Output("staleDays")]
        public Output<int> StaleDays { get; private set; } = null!;

        [Output("waitForCi")]
        public Output<bool?> WaitForCi { get; private set; } = null!;


        /// <summary>
        /// Create a AppConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppConfig(string name, AppConfigArgs args, CustomResourceOptions? options = null)
            : base("heroku:review/appConfig:AppConfig", name, args ?? new AppConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppConfig(string name, Input<string> id, AppConfigState? state = null, CustomResourceOptions? options = null)
            : base("heroku:review/appConfig:AppConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppConfig Get(string name, Input<string> id, AppConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AppConfig(name, id, state, options);
        }
    }

    public sealed class AppConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("automaticReviewApps")]
        public Input<bool>? AutomaticReviewApps { get; set; }

        [Input("baseName")]
        public Input<string>? BaseName { get; set; }

        [Input("deployTarget", required: true)]
        public Input<Inputs.AppConfigDeployTargetArgs> DeployTarget { get; set; } = null!;

        [Input("destroyStaleApps")]
        public Input<bool>? DestroyStaleApps { get; set; }

        [Input("orgRepo", required: true)]
        public Input<string> OrgRepo { get; set; } = null!;

        [Input("pipelineId", required: true)]
        public Input<string> PipelineId { get; set; } = null!;

        [Input("staleDays")]
        public Input<int>? StaleDays { get; set; }

        [Input("waitForCi")]
        public Input<bool>? WaitForCi { get; set; }

        public AppConfigArgs()
        {
        }
        public static new AppConfigArgs Empty => new AppConfigArgs();
    }

    public sealed class AppConfigState : global::Pulumi.ResourceArgs
    {
        [Input("automaticReviewApps")]
        public Input<bool>? AutomaticReviewApps { get; set; }

        [Input("baseName")]
        public Input<string>? BaseName { get; set; }

        [Input("deployTarget")]
        public Input<Inputs.AppConfigDeployTargetGetArgs>? DeployTarget { get; set; }

        [Input("destroyStaleApps")]
        public Input<bool>? DestroyStaleApps { get; set; }

        [Input("orgRepo")]
        public Input<string>? OrgRepo { get; set; }

        [Input("pipelineId")]
        public Input<string>? PipelineId { get; set; }

        [Input("repoId")]
        public Input<int>? RepoId { get; set; }

        [Input("staleDays")]
        public Input<int>? StaleDays { get; set; }

        [Input("waitForCi")]
        public Input<bool>? WaitForCi { get; set; }

        public AppConfigState()
        {
        }
        public static new AppConfigState Empty => new AppConfigState();
    }
}
