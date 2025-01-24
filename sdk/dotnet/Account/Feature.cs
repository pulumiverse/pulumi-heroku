// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Account
{
    /// <summary>
    /// This resource is used to create and manage [User Features](https://devcenter.heroku.com/articles/heroku-beta-features) on Heroku.
    /// 
    /// &gt; **NOTE:** If this resource's HCL is removed from a `.tf` file, the behavior is to disable account feature
    /// and remove resource from state.
    /// 
    /// ## Available Features
    /// 
    /// For a list of available features, use the [`heroku labs`](https://devcenter.heroku.com/articles/heroku-cli)
    /// command to fetch them for the current authenticated user.
    /// 
    /// The output will contain **User Features** that may be managed with this resource.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Existing account features can be imported using a combination of the account email (the email address tied to the Heroku API key)
    /// and the feature name.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import heroku:account/feature:Feature example_metrics name@example.com:metrics-request-volume
    /// ```
    /// </summary>
    [HerokuResourceType("heroku:account/feature:Feature")]
    public partial class Feature : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of account feature
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enable or disable the account feature
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Name of the account feature
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State of account feature
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Feature resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feature(string name, FeatureArgs args, CustomResourceOptions? options = null)
            : base("heroku:account/feature:Feature", name, args ?? new FeatureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feature(string name, Input<string> id, FeatureState? state = null, CustomResourceOptions? options = null)
            : base("heroku:account/feature:Feature", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Feature resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feature Get(string name, Input<string> id, FeatureState? state = null, CustomResourceOptions? options = null)
        {
            return new Feature(name, id, state, options);
        }
    }

    public sealed class FeatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the account feature
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Name of the account feature
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FeatureArgs()
        {
        }
        public static new FeatureArgs Empty => new FeatureArgs();
    }

    public sealed class FeatureState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of account feature
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable or disable the account feature
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the account feature
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// State of account feature
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public FeatureState()
        {
        }
        public static new FeatureState Empty => new FeatureState();
    }
}
