// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Pipeline
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Pipelines can be imported using the Pipeline `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import heroku:pipeline/pipeline:Pipeline foobar 12345678
    /// ```
    /// </summary>
    [HerokuResourceType("heroku:pipeline/pipeline:Pipeline")]
    public partial class Pipeline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the pipeline.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the pipeline. This block as the following required attributes:
        /// </summary>
        [Output("owner")]
        public Output<Outputs.PipelineOwner> Owner { get; private set; } = null!;


        /// <summary>
        /// Create a Pipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipeline(string name, PipelineArgs? args = null, CustomResourceOptions? options = null)
            : base("heroku:pipeline/pipeline:Pipeline", name, args ?? new PipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipeline(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
            : base("heroku:pipeline/pipeline:Pipeline", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Pipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipeline Get(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
        {
            return new Pipeline(name, id, state, options);
        }
    }

    public sealed class PipelineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the pipeline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the pipeline. This block as the following required attributes:
        /// </summary>
        [Input("owner")]
        public Input<Inputs.PipelineOwnerArgs>? Owner { get; set; }

        public PipelineArgs()
        {
        }
        public static new PipelineArgs Empty => new PipelineArgs();
    }

    public sealed class PipelineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the pipeline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the pipeline. This block as the following required attributes:
        /// </summary>
        [Input("owner")]
        public Input<Inputs.PipelineOwnerGetArgs>? Owner { get; set; }

        public PipelineState()
        {
        }
        public static new PipelineState Empty => new PipelineState();
    }
}
