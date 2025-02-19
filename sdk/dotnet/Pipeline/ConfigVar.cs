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
    /// Provides a resource to manage a pipeline's config vars.
    /// 
    /// The pipeline config var API can only retrieve config vars that can be set at the pipeline level.
    /// Additionally, these two supported pipeline stages are:
    /// - [Heroku CI](https://devcenter.heroku.com/articles/heroku-ci#setting-environment-variables-the-env-key) config vars (test stage)
    /// - [Review Apps](https://devcenter.heroku.com/articles/github-integration-review-apps#configuration) config vars (review stage)
    /// 
    /// The development, staging &amp; production stages do not have stage-level config vars, only those on the apps within each stage.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// This resource defines two config var attributes with one of them used for masking any sensitive/secret variables
    /// 
    /// during a `pulumi preview|apply` in a CI build, terminal, etc. This 'sensitive' distinction for config vars is unique to
    /// 
    /// this provider and not a built-in feature of the Heroku Platform API. Therefore, it will not be possible to import
    /// 
    /// this resource.
    /// 
    /// However, it is safe to define the resource in your configuration file and execute a `pulumi up`
    /// 
    /// as the end result is `noop` when the config vars already exist on the remote resource.
    /// </summary>
    [HerokuResourceType("heroku:pipeline/configVar:ConfigVar")]
    public partial class ConfigVar : global::Pulumi.CustomResource
    {
        /// <summary>
        /// All vars of a pipeline stage. This is marked `sensitive` so that `sensitive_vars` do not leak in the console/logs.
        /// </summary>
        [Output("allVars")]
        public Output<ImmutableDictionary<string, string>> AllVars { get; private set; } = null!;

        /// <summary>
        /// The UUID of an existing pipeline.
        /// </summary>
        [Output("pipelineId")]
        public Output<string> PipelineId { get; private set; } = null!;

        /// <summary>
        /// The pipeline's stage. Supported values are `test` &amp; `review`.
        /// </summary>
        [Output("pipelineStage")]
        public Output<string> PipelineStage { get; private set; } = null!;

        [Output("sensitiveVars")]
        public Output<ImmutableDictionary<string, string>?> SensitiveVars { get; private set; } = null!;

        /// <summary>
        /// Map of config vars that can be output in plaintext.
        /// </summary>
        [Output("vars")]
        public Output<ImmutableDictionary<string, string>?> Vars { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigVar resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigVar(string name, ConfigVarArgs args, CustomResourceOptions? options = null)
            : base("heroku:pipeline/configVar:ConfigVar", name, args ?? new ConfigVarArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigVar(string name, Input<string> id, ConfigVarState? state = null, CustomResourceOptions? options = null)
            : base("heroku:pipeline/configVar:ConfigVar", name, state, MakeResourceOptions(options, id))
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
                    "allVars",
                    "sensitiveVars",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigVar resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigVar Get(string name, Input<string> id, ConfigVarState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigVar(name, id, state, options);
        }
    }

    public sealed class ConfigVarArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of an existing pipeline.
        /// </summary>
        [Input("pipelineId", required: true)]
        public Input<string> PipelineId { get; set; } = null!;

        /// <summary>
        /// The pipeline's stage. Supported values are `test` &amp; `review`.
        /// </summary>
        [Input("pipelineStage", required: true)]
        public Input<string> PipelineStage { get; set; } = null!;

        [Input("sensitiveVars")]
        private InputMap<string>? _sensitiveVars;
        public InputMap<string> SensitiveVars
        {
            get => _sensitiveVars ?? (_sensitiveVars = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _sensitiveVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("vars")]
        private InputMap<string>? _vars;

        /// <summary>
        /// Map of config vars that can be output in plaintext.
        /// </summary>
        public InputMap<string> Vars
        {
            get => _vars ?? (_vars = new InputMap<string>());
            set => _vars = value;
        }

        public ConfigVarArgs()
        {
        }
        public static new ConfigVarArgs Empty => new ConfigVarArgs();
    }

    public sealed class ConfigVarState : global::Pulumi.ResourceArgs
    {
        [Input("allVars")]
        private InputMap<string>? _allVars;

        /// <summary>
        /// All vars of a pipeline stage. This is marked `sensitive` so that `sensitive_vars` do not leak in the console/logs.
        /// </summary>
        public InputMap<string> AllVars
        {
            get => _allVars ?? (_allVars = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _allVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The UUID of an existing pipeline.
        /// </summary>
        [Input("pipelineId")]
        public Input<string>? PipelineId { get; set; }

        /// <summary>
        /// The pipeline's stage. Supported values are `test` &amp; `review`.
        /// </summary>
        [Input("pipelineStage")]
        public Input<string>? PipelineStage { get; set; }

        [Input("sensitiveVars")]
        private InputMap<string>? _sensitiveVars;
        public InputMap<string> SensitiveVars
        {
            get => _sensitiveVars ?? (_sensitiveVars = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _sensitiveVars = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("vars")]
        private InputMap<string>? _vars;

        /// <summary>
        /// Map of config vars that can be output in plaintext.
        /// </summary>
        public InputMap<string> Vars
        {
            get => _vars ?? (_vars = new InputMap<string>());
            set => _vars = value;
        }

        public ConfigVarState()
        {
        }
        public static new ConfigVarState Empty => new ConfigVarState();
    }
}
