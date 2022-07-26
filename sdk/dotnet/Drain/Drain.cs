// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Drain
{
    [HerokuResourceType("heroku:drain/drain:Drain")]
    public partial class Drain : global::Pulumi.CustomResource
    {
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("sensitiveUrl")]
        public Output<string?> SensitiveUrl { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Drain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Drain(string name, DrainArgs args, CustomResourceOptions? options = null)
            : base("heroku:drain/drain:Drain", name, args ?? new DrainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Drain(string name, Input<string> id, DrainState? state = null, CustomResourceOptions? options = null)
            : base("heroku:drain/drain:Drain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "sensitiveUrl",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Drain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Drain Get(string name, Input<string> id, DrainState? state = null, CustomResourceOptions? options = null)
        {
            return new Drain(name, id, state, options);
        }
    }

    public sealed class DrainArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("sensitiveUrl")]
        private Input<string>? _sensitiveUrl;
        public Input<string>? SensitiveUrl
        {
            get => _sensitiveUrl;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sensitiveUrl = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public DrainArgs()
        {
        }
        public static new DrainArgs Empty => new DrainArgs();
    }

    public sealed class DrainState : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("sensitiveUrl")]
        private Input<string>? _sensitiveUrl;
        public Input<string>? SensitiveUrl
        {
            get => _sensitiveUrl;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sensitiveUrl = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("token")]
        public Input<string>? Token { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public DrainState()
        {
        }
        public static new DrainState Empty => new DrainState();
    }
}
