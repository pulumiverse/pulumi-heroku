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
    public static class GetPipeline
    {
        /// <summary>
        /// Use this data source to get information about a Heroku Pipeline.
        /// 
        /// &gt; **NOTE:**
        /// This data source can only be used to fetch information regarding a pipeline that has apps already associated to it.
        /// This is a limitation in the Heroku Platform API where it is not possible to query a pipeline without apps by its name.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_pipeline" "foobar" {
        ///   name = "pipeline-test-123"
        /// }
        /// ```
        /// </summary>
        public static Task<GetPipelineResult> InvokeAsync(GetPipelineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPipelineResult>("heroku:pipeline/getPipeline:getPipeline", args ?? new GetPipelineArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Heroku Pipeline.
        /// 
        /// &gt; **NOTE:**
        /// This data source can only be used to fetch information regarding a pipeline that has apps already associated to it.
        /// This is a limitation in the Heroku Platform API where it is not possible to query a pipeline without apps by its name.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_pipeline" "foobar" {
        ///   name = "pipeline-test-123"
        /// }
        /// ```
        /// </summary>
        public static Output<GetPipelineResult> Invoke(GetPipelineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPipelineResult>("heroku:pipeline/getPipeline:getPipeline", args ?? new GetPipelineInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Heroku Pipeline.
        /// 
        /// &gt; **NOTE:**
        /// This data source can only be used to fetch information regarding a pipeline that has apps already associated to it.
        /// This is a limitation in the Heroku Platform API where it is not possible to query a pipeline without apps by its name.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_pipeline" "foobar" {
        ///   name = "pipeline-test-123"
        /// }
        /// ```
        /// </summary>
        public static Output<GetPipelineResult> Invoke(GetPipelineInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPipelineResult>("heroku:pipeline/getPipeline:getPipeline", args ?? new GetPipelineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPipelineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The pipeline name or ID. Empty pipelines can only be retrieved by ID (UUID).
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetPipelineArgs()
        {
        }
        public static new GetPipelineArgs Empty => new GetPipelineArgs();
    }

    public sealed class GetPipelineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The pipeline name or ID. Empty pipelines can only be retrieved by ID (UUID).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetPipelineInvokeArgs()
        {
        }
        public static new GetPipelineInvokeArgs Empty => new GetPipelineInvokeArgs();
    }


    [OutputType]
    public sealed class GetPipelineResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The pipeline owner's ID
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The pipeline owner's type
        /// </summary>
        public readonly string OwnerType;

        [OutputConstructor]
        private GetPipelineResult(
            string id,

            string name,

            string ownerId,

            string ownerType)
        {
            Id = id;
            Name = name;
            OwnerId = ownerId;
            OwnerType = ownerType;
        }
    }
}
