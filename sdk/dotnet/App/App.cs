// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.App
{
    [HerokuResourceType("heroku:app/app:App")]
    public partial class App : global::Pulumi.CustomResource
    {
        [Output("acm")]
        public Output<bool> Acm { get; private set; } = null!;

        [Output("allConfigVars")]
        public Output<ImmutableDictionary<string, object>> AllConfigVars { get; private set; } = null!;

        [Output("buildpacks")]
        public Output<ImmutableArray<string>> Buildpacks { get; private set; } = null!;

        [Output("configVars")]
        public Output<ImmutableDictionary<string, object>> ConfigVars { get; private set; } = null!;

        [Output("gitUrl")]
        public Output<string> GitUrl { get; private set; } = null!;

        [Output("herokuHostname")]
        public Output<string> HerokuHostname { get; private set; } = null!;

        [Output("internalRouting")]
        public Output<bool> InternalRouting { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organization")]
        public Output<Outputs.AppOrganization?> Organization { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("sensitiveConfigVars")]
        public Output<ImmutableDictionary<string, string>> SensitiveConfigVars { get; private set; } = null!;

        [Output("space")]
        public Output<string?> Space { get; private set; } = null!;

        [Output("stack")]
        public Output<string> Stack { get; private set; } = null!;

        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        [Output("webUrl")]
        public Output<string> WebUrl { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("heroku:app/app:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
            : base("heroku:app/app:App", name, state, MakeResourceOptions(options, id))
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
                    "allConfigVars",
                    "sensitiveConfigVars",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
        {
            return new App(name, id, state, options);
        }
    }

    public sealed class AppArgs : global::Pulumi.ResourceArgs
    {
        [Input("acm")]
        public Input<bool>? Acm { get; set; }

        [Input("buildpacks")]
        private InputList<string>? _buildpacks;
        public InputList<string> Buildpacks
        {
            get => _buildpacks ?? (_buildpacks = new InputList<string>());
            set => _buildpacks = value;
        }

        [Input("configVars")]
        private InputMap<object>? _configVars;
        public InputMap<object> ConfigVars
        {
            get => _configVars ?? (_configVars = new InputMap<object>());
            set => _configVars = value;
        }

        [Input("internalRouting")]
        public Input<bool>? InternalRouting { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organization")]
        public Input<Inputs.AppOrganizationArgs>? Organization { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("sensitiveConfigVars")]
        private InputMap<string>? _sensitiveConfigVars;
        public InputMap<string> SensitiveConfigVars
        {
            get => _sensitiveConfigVars ?? (_sensitiveConfigVars = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _sensitiveConfigVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("space")]
        public Input<string>? Space { get; set; }

        [Input("stack")]
        public Input<string>? Stack { get; set; }

        public AppArgs()
        {
        }
        public static new AppArgs Empty => new AppArgs();
    }

    public sealed class AppState : global::Pulumi.ResourceArgs
    {
        [Input("acm")]
        public Input<bool>? Acm { get; set; }

        [Input("allConfigVars")]
        private InputMap<object>? _allConfigVars;
        public InputMap<object> AllConfigVars
        {
            get => _allConfigVars ?? (_allConfigVars = new InputMap<object>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, object>());
                _allConfigVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("buildpacks")]
        private InputList<string>? _buildpacks;
        public InputList<string> Buildpacks
        {
            get => _buildpacks ?? (_buildpacks = new InputList<string>());
            set => _buildpacks = value;
        }

        [Input("configVars")]
        private InputMap<object>? _configVars;
        public InputMap<object> ConfigVars
        {
            get => _configVars ?? (_configVars = new InputMap<object>());
            set => _configVars = value;
        }

        [Input("gitUrl")]
        public Input<string>? GitUrl { get; set; }

        [Input("herokuHostname")]
        public Input<string>? HerokuHostname { get; set; }

        [Input("internalRouting")]
        public Input<bool>? InternalRouting { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organization")]
        public Input<Inputs.AppOrganizationGetArgs>? Organization { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sensitiveConfigVars")]
        private InputMap<string>? _sensitiveConfigVars;
        public InputMap<string> SensitiveConfigVars
        {
            get => _sensitiveConfigVars ?? (_sensitiveConfigVars = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _sensitiveConfigVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("space")]
        public Input<string>? Space { get; set; }

        [Input("stack")]
        public Input<string>? Stack { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        [Input("webUrl")]
        public Input<string>? WebUrl { get; set; }

        public AppState()
        {
        }
        public static new AppState Empty => new AppState();
    }
}
