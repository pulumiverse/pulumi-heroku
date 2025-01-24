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
    /// <summary>
    /// Provides a resource for managing [inbound rulesets](https://devcenter.heroku.com/articles/platform-api-reference#inbound-ruleset) for Heroku Private Spaces.
    /// 
    /// ## Example Usage
    /// </summary>
    [HerokuResourceType("heroku:space/inboundRuleset:InboundRuleset")]
    public partial class InboundRuleset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// At least one `rule` block. Rules are documented below.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.InboundRulesetRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The ID of the space.
        /// </summary>
        [Output("space")]
        public Output<string> Space { get; private set; } = null!;


        /// <summary>
        /// Create a InboundRuleset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InboundRuleset(string name, InboundRulesetArgs args, CustomResourceOptions? options = null)
            : base("heroku:space/inboundRuleset:InboundRuleset", name, args ?? new InboundRulesetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InboundRuleset(string name, Input<string> id, InboundRulesetState? state = null, CustomResourceOptions? options = null)
            : base("heroku:space/inboundRuleset:InboundRuleset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InboundRuleset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InboundRuleset Get(string name, Input<string> id, InboundRulesetState? state = null, CustomResourceOptions? options = null)
        {
            return new InboundRuleset(name, id, state, options);
        }
    }

    public sealed class InboundRulesetArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.InboundRulesetRuleArgs>? _rules;

        /// <summary>
        /// At least one `rule` block. Rules are documented below.
        /// </summary>
        public InputList<Inputs.InboundRulesetRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.InboundRulesetRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The ID of the space.
        /// </summary>
        [Input("space", required: true)]
        public Input<string> Space { get; set; } = null!;

        public InboundRulesetArgs()
        {
        }
        public static new InboundRulesetArgs Empty => new InboundRulesetArgs();
    }

    public sealed class InboundRulesetState : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.InboundRulesetRuleGetArgs>? _rules;

        /// <summary>
        /// At least one `rule` block. Rules are documented below.
        /// </summary>
        public InputList<Inputs.InboundRulesetRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.InboundRulesetRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The ID of the space.
        /// </summary>
        [Input("space")]
        public Input<string>? Space { get; set; }

        public InboundRulesetState()
        {
        }
        public static new InboundRulesetState Empty => new InboundRulesetState();
    }
}
