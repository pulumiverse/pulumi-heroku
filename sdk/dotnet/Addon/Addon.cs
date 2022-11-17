// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Addon
{
    [HerokuResourceType("heroku:addon/addon:Addon")]
    public partial class Addon : global::Pulumi.CustomResource
    {
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("config")]
        public Output<ImmutableDictionary<string, object>?> Config { get; private set; } = null!;

        [Output("configVarValues")]
        public Output<ImmutableDictionary<string, string>> ConfigVarValues { get; private set; } = null!;

        [Output("configVars")]
        public Output<ImmutableArray<string>> ConfigVars { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("providerId")]
        public Output<string> ProviderId { get; private set; } = null!;


        /// <summary>
        /// Create a Addon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Addon(string name, AddonArgs args, CustomResourceOptions? options = null)
            : base("heroku:addon/addon:Addon", name, args ?? new AddonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Addon(string name, Input<string> id, AddonState? state = null, CustomResourceOptions? options = null)
            : base("heroku:addon/addon:Addon", name, state, MakeResourceOptions(options, id))
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
                    "configVarValues",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Addon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Addon Get(string name, Input<string> id, AddonState? state = null, CustomResourceOptions? options = null)
        {
            return new Addon(name, id, state, options);
        }
    }

    public sealed class AddonArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("config")]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        public AddonArgs()
        {
        }
        public static new AddonArgs Empty => new AddonArgs();
    }

    public sealed class AddonState : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("config")]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        [Input("configVarValues")]
        private InputMap<string>? _configVarValues;
        public InputMap<string> ConfigVarValues
        {
            get => _configVarValues ?? (_configVarValues = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _configVarValues = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("configVars")]
        private InputList<string>? _configVars;
        public InputList<string> ConfigVars
        {
            get => _configVars ?? (_configVars = new InputList<string>());
            set => _configVars = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        public AddonState()
        {
        }
        public static new AddonState Empty => new AddonState();
    }
}
