// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Space
{
    [HerokuResourceType("heroku:space/space:Space")]
    public partial class Space : global::Pulumi.CustomResource
    {
        [Output("cidr")]
        public Output<string?> Cidr { get; private set; } = null!;

        [Output("dataCidr")]
        public Output<string> DataCidr { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        [Output("outboundIps")]
        public Output<ImmutableArray<string>> OutboundIps { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("shield")]
        public Output<bool?> Shield { get; private set; } = null!;


        /// <summary>
        /// Create a Space resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Space(string name, SpaceArgs args, CustomResourceOptions? options = null)
            : base("heroku:space/space:Space", name, args ?? new SpaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Space(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
            : base("heroku:space/space:Space", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Space resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Space Get(string name, Input<string> id, SpaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Space(name, id, state, options);
        }
    }

    public sealed class SpaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("dataCidr")]
        public Input<string>? DataCidr { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organization", required: true)]
        public Input<string> Organization { get; set; } = null!;

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shield")]
        public Input<bool>? Shield { get; set; }

        public SpaceArgs()
        {
        }
        public static new SpaceArgs Empty => new SpaceArgs();
    }

    public sealed class SpaceState : global::Pulumi.ResourceArgs
    {
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        [Input("dataCidr")]
        public Input<string>? DataCidr { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("outboundIps")]
        private InputList<string>? _outboundIps;
        public InputList<string> OutboundIps
        {
            get => _outboundIps ?? (_outboundIps = new InputList<string>());
            set => _outboundIps = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shield")]
        public Input<bool>? Shield { get; set; }

        public SpaceState()
        {
        }
        public static new SpaceState Empty => new SpaceState();
    }
}
